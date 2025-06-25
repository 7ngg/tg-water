package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

type AiResponse struct {
	TotalCalories float64 `json:"total_calories"`
	FoodItems     []struct {
		Name      string `json:"name"`
		Quantity  string `json:"quantity"`
		Nutrition struct {
			Calories float64 `json:"calories"`
			Protein  float64 `json:"protein"`
			Carbs    float64 `json:"carbs"`
			Fat      float64 `json:"fat"`
			Fiber    float64 `json:"fiber"`
			Sugar    float64 `json:"sugar"`
		} `json:"nutrition"`
	} `json:"food_items"`
	Analyze string      `json:"analyze"`
	Message any `json:"message"`
}

type AiService struct {
	HttpClient *http.Client
}

func New() *AiService {
	return &AiService{
		HttpClient: &http.Client{},
	}
}

func (s *AiService) AnalyzeImage(ctx context.Context, path string) (AiResponse, error) {
	url := "https://ai-calorie-tracker.p.rapidapi.com/analyze"

	fileData, err := os.ReadFile(path)
	if err != nil {
		return AiResponse{}, fmt.Errorf("failed to read file: %w", err)
	}

	contentType := http.DetectContentType(fileData)

	if !strings.HasPrefix(contentType, "image/") {
		return AiResponse{}, fmt.Errorf("file is not an image, detected type: %s", contentType)
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="image"; filename="image.jpg"`)
	h.Set("Content-Type", "image/jpeg")

	fileWriter, err := writer.CreatePart(h)
	if err != nil {
		return AiResponse{}, fmt.Errorf("failed to create form part: %w", err)
	}

	_, err = fileWriter.Write(fileData)
	if err != nil {
		return AiResponse{}, fmt.Errorf("failed to write file data: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return AiResponse{}, fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return AiResponse{}, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-RapidAPI-Key", "9e977b0f6dmsh720fb50431be912p105275jsn018081b5c3e7")
	req.Header.Set("X-RapidAPI-Host", "ai-calorie-tracker.p.rapidapi.com")
	req.Header.Set("Accept", "application/json")

	res, err := s.HttpClient.Do(req)
	if err != nil {
		return AiResponse{}, err
	}
	defer res.Body.Close()

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return AiResponse{}, err
	}

	var data AiResponse
	err = json.Unmarshal(responseBytes, &data)
	if err != nil {
		return AiResponse{}, fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(responseBytes))
	}

	return data, nil
}

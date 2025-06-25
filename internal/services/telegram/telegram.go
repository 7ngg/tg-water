package telegram

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-telegram/bot"
)

type TelegramService struct {
	bot        *bot.Bot
	httpClient *http.Client
}

func New(b *bot.Bot) *TelegramService {
	return &TelegramService{
		bot:        b,
		httpClient: &http.Client{},
	}
}

func (s *TelegramService) DownloadPhoto(ctx context.Context, fileID string) (string, error) {
	metadata, err := s.bot.GetFile(ctx, &bot.GetFileParams{FileID: fileID})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", s.bot.Token(), metadata.FilePath)
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fileName := fmt.Sprintf("tmp/%s.jpg", metadata.FileID)
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return "", err
	}

	return fileName, nil
}

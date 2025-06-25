package tgbot

import (
	"fmt"

	"github.com/7ngg/trackly/internal/services/ai"
)

func GetNutritionReponse(res ai.AiResponse) string {
	protein := 0.0
	carbs := 0.0
	fat := 0.0
	fiber := 0.0
	sugar := 0.0

	for _, f := range res.FoodItems {
		protein += f.Nutrition.Protein
		carbs += f.Nutrition.Carbs
		fat += f.Nutrition.Fat
		fiber += f.Nutrition.Fiber
		sugar += f.Nutrition.Sugar
	}

	html := fmt.Sprintf(
`<b>🍽️ Nutrition Facts for Your Meal</b>
<pre>Calories: %6.1f
	Protein:  %6.1fg
	Carbs:    %6.1fg
	Fat:      %6.1fg
	Fiber:    %6.1fg
	Sugar:    %6.1fg</pre>
<b>🤖 AI Analysis:</b> %s
	`, res.TotalCalories, protein, carbs, fat, fiber, sugar, res.Analyze)

	return html
}

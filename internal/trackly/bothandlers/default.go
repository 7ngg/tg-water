package bothandlers

import (
	"context"
	"log/slog"

	"github.com/7ngg/trackly/internal/lib/logger"
	"github.com/7ngg/trackly/internal/lib/tgbot"
	"github.com/7ngg/trackly/internal/services/ai"
	"github.com/7ngg/trackly/internal/services/telegram"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(log *slog.Logger) func(ctx context.Context, b *bot.Bot, u *models.Update) {
	return func(ctx context.Context, b *bot.Bot, u *models.Update) {
		if len(u.Message.Photo) == 0 {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "I can't recognize what you've typed... Please, check commands list",
			})

			return
		}

		tg := telegram.New(b)
		ai := ai.New()

		photo := u.Message.Photo[len(u.Message.Photo)-1]

		file, err := tg.DownloadPhoto(ctx, photo.FileID)
		if err != nil {
			log.Error("failed to download photo", logger.Err(err))
			tgbot.RespondWithError(ctx, b, int(u.Message.Chat.ID))
			return
		}

		some, err := ai.AnalyzeImage(ctx, file)
		if err != nil {
			log.Error("failed to analyze image", logger.Err(err))
			tgbot.RespondWithError(ctx, b, int(u.Message.Chat.ID))
			return
		}

		msg := tgbot.GetNutritionReponse(some)
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    u.Message.Chat.ID,
			Text:      msg,
			ParseMode: models.ParseModeHTML,
		})

		if err != nil {
			log.Error("failed to send message", logger.Err(err))
		}
	}
}

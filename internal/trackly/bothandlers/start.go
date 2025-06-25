package bothandlers

import (
	"context"

	"github.com/7ngg/trackly/internal/storage/sqlite"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type UserStorage interface {
	GetByTgID(ctx context.Context, telegramID int64) (sqlite.User, error)
	SaveUser(ctx context.Context, arg sqlite.SaveUserParams) (sqlite.User, error)
}

func Start(storage UserStorage) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, u *models.Update) {
		chatID := u.Message.Chat.ID
		telegramID := u.Message.From.ID

		existing, _ := storage.GetByTgID(ctx, telegramID)

		if existing.TelegramID != 0 {
			return
		}

		_, err := storage.SaveUser(ctx, sqlite.SaveUserParams{
			TelegramID: telegramID,
			ChatID:     chatID,
		})

		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "Unable to add user... Please, send '/start' command once again",
			})
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "You're all set! Welcome to the club",
		})
	}
}

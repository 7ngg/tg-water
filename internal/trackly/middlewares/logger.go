package middlewares

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Logger(log *slog.Logger) func(bot.HandlerFunc) bot.HandlerFunc {
	return func(next bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, b *bot.Bot, u *models.Update) {
			log.Info("incoming command",
				slog.Int64("telegram_id", u.Message.From.ID),
				slog.Int64("chat_id", u.Message.Chat.ID),
				slog.String("command", u.Message.Text))
			next(ctx, b, u)
		}
	}
}

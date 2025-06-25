package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
)

func RespondWithError(ctx context.Context, b *bot.Bot, chatID int) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text: "Looks like something went wrong... Please, try again",
	})
}

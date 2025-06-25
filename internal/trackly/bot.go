package trackly

import (
	"log/slog"

	"github.com/7ngg/trackly/internal/storage/sqlite"
	"github.com/7ngg/trackly/internal/trackly/bothandlers"
	"github.com/7ngg/trackly/internal/trackly/middlewares"
	"github.com/go-telegram/bot"
)

type TracklyBot struct {
	Bot     *bot.Bot
	Storage *sqlite.Storage
}

func New(token string, storage *sqlite.Storage, log *slog.Logger) (*TracklyBot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(bothandlers.DefaultHandler(log)),
		bot.WithMiddlewares(middlewares.Logger(log)),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, bothandlers.Start(storage.DB))

	trackly := &TracklyBot{
		Bot:     b,
		Storage: storage,
	}

	return trackly, err
}

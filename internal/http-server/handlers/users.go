package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/7ngg/trackly/internal/lib/api"
	"github.com/7ngg/trackly/internal/storage/sqlite"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type response struct {
	api.ReponseBase
	Users []sqlite.User `json:"users"`
}

type UsersGetter interface {
	ListUsers(ctx context.Context) ([]sqlite.User, error)
}

func GetAllUsers(log *slog.Logger, getter UsersGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "handlers.users.get-all.New"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		users, err := getter.ListUsers(r.Context())
		if err != nil {
			render.JSON(w, r, api.Error(err.Error(), http.StatusInternalServerError))
			return
		}

		render.JSON(w, r, response{
			ReponseBase: *api.Ok(http.StatusOK),
			Users:       users,
		})
	}
}

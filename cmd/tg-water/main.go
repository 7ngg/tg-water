package main

import (
	"fmt"

	"github.com/7ngg/tg-water/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi router"

	// TODO: run server
}

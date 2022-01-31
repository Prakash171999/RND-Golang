package main

import (
	"fx-example/handler"
	"fx-example/logger"
	"fx-example/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		server.Module,
		logger.Module,
		fx.Provide(
			handler.NewHandler,
		),
	)
	app.Run()
}

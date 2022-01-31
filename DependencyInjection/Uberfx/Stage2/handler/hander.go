package handler

import (
	"fx-example/logger"
	"io"
	"net/http"
)

func NewHandler(logger logger.Logger) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Handler Called")
		io.WriteString(w, "Hello world!")
	})
}

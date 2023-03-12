package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)

	return mux
}

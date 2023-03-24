package main

import (
	"fmt"
	"net/http"
)

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "world")
}

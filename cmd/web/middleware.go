package main

import "net/http"

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

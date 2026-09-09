package handlers

import "net/http"

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	// This handler will create a panic
	panic("Boom 💥")
}

package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h1>Halo nama saya Afif!</h1>")
	}
}

package middleware

import (
	"log"
	"net/http"
)

func Testing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("This is a testing middleware .... !")
		next.ServeHTTP(w, r)
	})
}

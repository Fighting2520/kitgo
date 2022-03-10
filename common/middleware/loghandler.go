package middleware

import (
	"io"
	"net/http"
)

func LoggerHandler(writer io.Writer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {

			}()
			next.ServeHTTP(w, r)
		})
	}
}

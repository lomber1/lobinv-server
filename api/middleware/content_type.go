package middleware

import "net/http"

func contentType(contentType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		next.ServeHTTP(w, r)
	})
}

func ContentTypeText(next http.Handler) http.Handler {
	return contentType("text/plain; charset=utf-8", next)
}

func ContentTypeJSON(next http.Handler) http.Handler {
	return contentType("application/json; charset=utf-8", next)
}

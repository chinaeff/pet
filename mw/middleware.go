package mw

import "net/http"

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized. Token is required.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

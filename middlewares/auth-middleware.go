package middlewares

import (
	"context"
	"net/http"
)

type contextKey string
const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")

		if err != nil{
			next.ServeHTTP(w, r)
			return 
		}

		_ = cookie.Value
		userID := 1

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
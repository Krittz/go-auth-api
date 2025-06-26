package middleware

import (
	"context"
	"go-auth-api/pkg/config"
	//"log"
	"net/http"
	//"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "userID"

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Token não encontrado", http.StatusUnauthorized)
			return
		}
		cfg := config.LoadConfig()
		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Claims inválidos", http.StatusUnauthorized)
			return
		}
		sub, ok := claims["sub"].(float64)
		if !ok {
			http.Error(w, "ID inválido no token", http.StatusUnauthorized)
			return
		}
		userID := int64(sub)

		// Injeta o userID no context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDKey).(int64)
	return id, ok
}

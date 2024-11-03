package middleware

import (
	"net/http"
	"strings"

	"course-management-api/auth"
)

func VerificarToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := auth.ValidarToken(tokenString)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		if usuarioID, ok := claims["usuario_id"].(string); ok {
			r.Header.Set("usuario_id", usuarioID)
		} else {
			http.Error(w, "ID de usuário inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

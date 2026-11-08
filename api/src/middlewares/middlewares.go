package middlewares

import (
	"api/src/auth"
	"api/src/response"
	"fmt"
	"log/slog"
	"net/http"
)

// Logger registra informações sobre a requisição HTTP e, em seguida,
// encaminha a requisição para o próximo handler.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Host))
		next(w, r)
	}
}

// Auth faz a verificacao se o usuário está autenticado na requisicao
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			slog.Warn("Não autorizado.", "warn", err)
			response.ResponseError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

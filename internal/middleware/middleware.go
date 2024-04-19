package middleware

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/zakhaev26/recruitments/utils"
)

func AuthorizationMiddleware(next http.HandlerFunc, allowedRoles []string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if len(allowedRoles) == 1 && allowedRoles[0] == "*" {
			next(w, r)
			return
		}

		token, err := utils.ExtractTokenFromHeader(r)
		if err != nil {
			log.Error(err)
			http.Error(w, "Failed to extract token from headers", http.StatusBadRequest)
			return
		}

		claims, err := utils.ParseAccessToken(token)
		if err != nil {
			log.Error(err)
			http.Error(w, "Failed to parse token from headers", http.StatusBadRequest)
			return
		}

		for _, v := range allowedRoles {
			if v == claims.UserType {
				next(w, r)
				return
			}
		}

		log.Error(err)
		http.Error(w, "Failed to find user", http.StatusInternalServerError)
	}
}

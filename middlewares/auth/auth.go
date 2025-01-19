package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

type userNameKey struct{}

const (
	tokenPrefix = "UT"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		userName, err := validateToken(token)
		if err != nil {
			log.Println(err)
			http.Error(w, `{"reason": "invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userNameKey{}, userName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(token string) (string, error) {
	tElems := strings.SplitN(token, "_", 2)
	if len(tElems) < 2 {
		return "", errors.New("invalid token")
	}

	tType, tUserName := tElems[0], tElems[1]
	if tType != tokenPrefix {
		return "", errors.New("invalid token")
	}
	return tUserName, nil
}

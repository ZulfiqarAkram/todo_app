package middleware

import (
	"context"
	"fmt"
	"net/http"
	"todo_app/auth"
)

type contextKey int

const AuthenticatedUserKey contextKey = 0

var AllowedPathWithToken = []string{
	"/api/login",
	"/api/register",
}

type Authorization struct {
	handler http.Handler
}

func (l *Authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	for _, path := range AllowedPathWithToken {
		if path == r.URL.Path {
			l.handler.ServeHTTP(w, r)
			return
		}
	}
	payLoad, err := IsValidToken(token)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return

	}
	contextWithUser := context.WithValue(r.Context(), AuthenticatedUserKey, payLoad)
	rWithUser := r.WithContext(contextWithUser)
	l.handler.ServeHTTP(w, rWithUser)
}

//Constructor
func NewAuthorization(handlerToWrap http.Handler) *Authorization {
	fmt.Println("Constructor NewAuthorization()")
	return &Authorization{handlerToWrap}
}

func IsValidToken(token string) (map[string]interface{}, error) {
	payLoadResult, err := auth.JWTManager.Decode(token)
	if err != nil {
		return payLoadResult, err
	}
	return payLoadResult, nil
}

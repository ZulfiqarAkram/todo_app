package middleware

import (
	"context"
	"net/http"
	"todo_app/auth"
)

type Authentication struct {
	Handler    http.Handler
	JWTManager *auth.JwtAuth
}
type contextKey int

const AuthenticatedUserKey contextKey = 0

var allowedPathWithToken = []string{
	"/api/v1/user/login",
	"/api/v1/user/register",
}

//Authentication Constructor
func New(handlerToWrap http.Handler, JWTManager *auth.JwtAuth) *Authentication {
	return &Authentication{handlerToWrap, JWTManager}
}

func (auth Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	for _, path := range allowedPathWithToken {
		if path == r.URL.Path {
			auth.Handler.ServeHTTP(w, r)
			return
		}
	}
	payload, err := auth.JWTManager.Decode(token)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return

	}
	contextWithUser := context.WithValue(r.Context(), AuthenticatedUserKey, payload)
	rWithUser := r.WithContext(contextWithUser)
	auth.Handler.ServeHTTP(w, rWithUser)
}

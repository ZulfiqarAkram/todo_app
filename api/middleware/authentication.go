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

var AllowedPathWithToken = []string{
	"/api/v1/user/login",
	"/api/v1/user/register",
}

//Authentication Constructor
func New(handlerToWrap http.Handler, JWTManager *auth.JwtAuth) *Authentication {
	return &Authentication{handlerToWrap, JWTManager}
}

func (a Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	for _, path := range AllowedPathWithToken {
		if path == r.URL.Path {
			a.Handler.ServeHTTP(w, r)
			return
		}
	}
	payLoad, err := a.JWTManager.Decode(token)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return

	}
	contextWithUser := context.WithValue(r.Context(), AuthenticatedUserKey, payLoad)
	rWithUser := r.WithContext(contextWithUser)
	a.Handler.ServeHTTP(w, rWithUser)
}

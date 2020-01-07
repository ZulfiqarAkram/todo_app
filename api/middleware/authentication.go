package middleware

import (
	"context"
	"net/http"
)

type Authentication struct {
	Handler http.Handler
}
type contextKey int

const AuthenticatedUserKey contextKey = 0

var AllowedPathWithToken = []string{
	"/api/v1/login",
	"/api/v1/register",
}

func (a Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	for _, path := range AllowedPathWithToken {
		if path == r.URL.Path {
			a.Handler.ServeHTTP(w, r)
			return
		}
	}
	payLoad, err := IsValidToken(token) //error here
	if err != nil {
		http.Error(w, err.Error(), 401)
		return

	}
	contextWithUser := context.WithValue(r.Context(), AuthenticatedUserKey, payLoad)
	rWithUser := r.WithContext(contextWithUser)
	a.Handler.ServeHTTP(w, rWithUser)
}

//Authentication Constructor
func New(handlerToWrap http.Handler) *Authentication {
	return &Authentication{handlerToWrap}
}

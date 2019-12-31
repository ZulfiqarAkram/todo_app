package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo_app/auth"
)

type contextKey int

const AuthenticatedUserKey contextKey = 0

type Authorization struct {
	handler http.Handler
}

func (l *Authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	// Pass Authorization for Login and Registration
	if token == "" {
		l.handler.ServeHTTP(w, r)
	} else {
		payLoad, err, isValid := IsValidToken(token)
		if isValid {
			contextWithUser := context.WithValue(r.Context(), AuthenticatedUserKey, payLoad)
			rWithUser := r.WithContext(contextWithUser)
			l.handler.ServeHTTP(w, rWithUser)
		} else {
			http.Error(w, err.Error(), 401)
			return
		}
		log.Printf("=> %s %s", r.Method, r.URL.Path)
	}
}

//Constructor
func NewAuthorization(handlerToWrap http.Handler) *Authorization {
	return &Authorization{handlerToWrap}
}

func IsValidToken(token string) (map[string]interface{}, error, bool) {
	payLoadResult, err := auth.JWTManager.Decode(token)
	if err != nil {
		fmt.Println(err)
		return payLoadResult, err, false
	}
	fmt.Println(payLoadResult)
	return payLoadResult, err, true
}

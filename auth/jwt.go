package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mzulfiqar10p/todo_app/setting"
	"time"
)

//JWT manages jwt
type JWT interface {
	Sign(data map[string]interface{}) (token string, err error)
	Decode(token string) (payload map[string]interface{}, err error)
}

type JwtAuth struct {
	secret    string
	expiresIn time.Duration
}

func CreateJWTManager() *JwtAuth {
	return NewJWTWithConf(setting.SecretKey, setting.Duration)
}

//NewJWTWithConf create new jwt manager
func NewJWTWithConf(secret string, expiry time.Duration) *JwtAuth {
	return &JwtAuth{secret: secret, expiresIn: expiry}
}

func (j *JwtAuth) Sign(payload map[string]interface{}) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(j.expiresIn).Unix()
	for k, v := range payload {
		claims[k] = v
	}
	return token.SignedString([]byte(j.secret))

}
func (j *JwtAuth) Decode(token string) (payload map[string]interface{}, err error) {

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("token is not valid anymore")
	}

	return t.Claims.(jwt.MapClaims), nil
}

package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	cfg "todo_app/config"
)

type jwtAuth struct {
	secret    string
	expiresIn time.Duration
}

var JWTManager JWT

func CreateJWTManager() {
	JWTManager = NewJWTWithConf(cfg.SecretKey, cfg.Duration)
}

//JWT manages jwt
type JWT interface {
	Sign(data map[string]interface{}) (token string, err error)
	Decode(token string) (payload map[string]interface{}, err error)
}

//NewJWTWithConf create new jwt manager
func NewJWTWithConf(secret string, expiry time.Duration) JWT {
	return &jwtAuth{secret: secret, expiresIn: expiry}
}

func (j *jwtAuth) Sign(payload map[string]interface{}) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(j.expiresIn).Unix()
	for k, v := range payload {
		claims[k] = v
	}
	return token.SignedString([]byte(j.secret))

}

func (j *jwtAuth) Decode(token string) (payload map[string]interface{}, err error) {

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

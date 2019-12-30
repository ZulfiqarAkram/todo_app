package auth

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestJWT_CanSign(t *testing.T) {

	secret := "123"
	duration := 1 * time.Second
	jwtManager := NewJWTWithConf(secret, duration)
	token, err := jwtManager.Sign(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, len(strings.Split(token, ".")), 3)
}

func TestJWT_ShouldDecodeBeforeOneSecond(t *testing.T) {
	secret := "123"
	duration := 1 * time.Second
	jwtManager := NewJWTWithConf(secret, duration)
	token, err := jwtManager.Sign(nil)
	assert.Nil(t, err)
	_, err = jwtManager.Decode(token)
	assert.Nil(t, err)
}

func TestJWT_ShouldExpireAfterOneSecond(t *testing.T) {
	secret := "123"
	duration := 1 * time.Second
	jwtManager := NewJWTWithConf(secret, duration)
	token, err := jwtManager.Sign(nil)
	time.Sleep(2 * time.Second)
	assert.Nil(t, err)
	_, err = jwtManager.Decode(token)
	assert.NotNil(t, err)
}

func TestJWT_ShouldGetPayloadAfterDecode(t *testing.T) {
	secret := "123"
	duration := 1 * time.Second
	jwtManager := NewJWTWithConf(secret, duration)
	x := map[string]interface{}{
		"firstName": "John",
		"lastName":  "Doe",
	}
	token, err := jwtManager.Sign(x)
	assert.Nil(t, err)
	payload, err := jwtManager.Decode(token)
	assert.Nil(t, err)

	for k, v := range x {
		assert.Equal(t, payload[k], v)
	}

}

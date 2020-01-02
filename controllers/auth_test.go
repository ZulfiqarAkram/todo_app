package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo_app/auth"
	"todo_app/middleware"
)

func TestRegister_ShouldAddNewUserIntoDB(t *testing.T) {
	newUser := []byte(`{"email_address":"moon@gmail.com","password":"123456","username":"moon_moon"}`)
	body, err := RegisterUser(newUser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var jsonRes boomErr
	err = json.NewDecoder(body).Decode(&jsonRes)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	expected := "New user has been registered."
	if jsonRes.Message != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body.String(), expected)
	}

}
func TestLogin_ShouldGenerateToken(t *testing.T) {
	//Register
	newUser := []byte(`{"email_address":"moon@gmail.com","password":"123456","username":"moon_moon"}`)
	_, err := RegisterUser(newUser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	//Login
	userCredentials := []byte(`{"email_address":"moon@gmail.com","password":"123456"}`)
	token, err := LoginUser(userCredentials)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.NotEmpty(t, token)
	assert.Equal(t, len(strings.Split(token, ".")), 3)
}
func TestLogin_ShouldThrowUnAuthorizeError(t *testing.T) {
	//Register
	newUser := []byte(`{"email_address":"moon@gmail.com","password":"123456","username":"moon_moon"}`)
	_, err := RegisterUser(newUser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	//Login
	invalidUserCredentials := []byte(`{"email_address":"moon123@gmail.com","password":"123444"}`)
	message, err := LoginUser(invalidUserCredentials)
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := "Unauthorized"
	assert.Equal(t, expected, message)
}

func RegisterUser(newUser []byte) (*bytes.Buffer, error) {
	// Register JWT
	auth.CreateJWTManager()

	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(newUser))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	//Middleware
	n := negroni.New()
	n.UseHandler(middleware.NewAuthorization(handler))
	n.ServeHTTP(rr, req)
	return rr.Body, nil

}
func LoginUser(userCredentials []byte) (string, error) {
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(userCredentials))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	//Middleware
	n := negroni.New()
	n.UseHandler(middleware.NewAuthorization(handler))
	n.ServeHTTP(rr, req)

	var jsonRes boomErr
	err = json.NewDecoder(rr.Body).Decode(&jsonRes)
	if err != nil {
		return "", err
	}

	fmt.Println(jsonRes)
	return jsonRes.Message, nil
}
func RegisterAndLoginUser() (string, error) {
	//Register
	newUser := []byte(`{""username":"moon_moon", "email_address":"moon@gmail.com","password":"123456"}`)
	_, err := RegisterUser(newUser)
	if err != nil {
		return "", err
	}
	//Login
	invalidUserCredentials := []byte(`{"email_address":"moon@gmail.com","password":"123456"}`)
	token, err := LoginUser(invalidUserCredentials)
	if err != nil {
		return "", err
	}
	return token, err
}

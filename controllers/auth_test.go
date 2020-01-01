package controllers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister_ShouldAddNewUserIntoDB(t *testing.T) {
	userCredentials := []byte(`{"email_address":"moon@gmail.com","password":"123456","username":"moon"}`)
	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(userCredentials))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	expected := "New user has been registered."
	if strings.Compare(expected, rr.Body.String()) == -1 {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}
func TestLogin_ShouldGenerateToken(t *testing.T) {
	userCredentials := []byte(`{"email_address":"ali@gmail.com","password":"123456"}`)
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(userCredentials))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	token := rr.Body.String()
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, len(strings.Split(token, ".")), 2)
}
func TestLogin_ShouldThrowUnAuthorizeError(t *testing.T) {
	userCredentials := []byte(`{"email_address":"ali123@gmail.com","password":"65465"}`)
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(userCredentials))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	expected := "Unauthorized\n"
	assert.Equal(t, expected, rr.Body.String())
}

package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRegister_ShouldAddNewUserIntoDB(t *testing.T) {
	var RegisterUrl = "http://localhost:8080/api/v1/user/register"
	newUser := []byte(`{
		"id":"0",
		"username" : "raja",
		"email_address" : "raja123@gmail.com",
		"password":"123456"
	}`)
	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	boomRes, err := SendRequest("POST", RegisterUrl, newUser, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(boomRes.Message)
	assert.Equal(t, boomRes.Message, "New user has been registered.")

}
func TestLogin_ShouldGenerateToken(t *testing.T) {
	//Register
	var RegisterUrl = "http://localhost:8080/api/v1/user/register"
	newUser := []byte(`{
		"id":"0",
		"username" : "hamza",
		"email_address" : "hamza123@gmail.com",
		"password":"123456"
	}`)
	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	boomRes, err := SendRequest("POST", RegisterUrl, newUser, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(boomRes.Message)
	assert.Equal(t, boomRes.Message, "New user has been registered.")

	//Login
	loginCredentials := []byte(`{"email_address":"hamza123@gmail.com","password":"123456"}`)
	loginUrl := "http://localhost:8080/api/v1/user/login"
	boomRes, err = SendRequest("POST", loginUrl, loginCredentials, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(boomRes.Message)
	token := boomRes.Message
	assert.Equal(t, len(strings.Split(token, ".")), 3)
}
func TestLogin_ShouldThrowUnAuthorizeError(t *testing.T) {
	//Register
	var RegisterUrl = "http://localhost:8080/api/v1/user/register"
	newUser := []byte(`{
		"id":"0",
		"username" : "kamran",
		"email_address" : "kamran123@gmail.com",
		"password":"123456"
	}`)
	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	boomRes, err := SendRequest("POST", RegisterUrl, newUser, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(boomRes.Message)
	assert.Equal(t, boomRes.Message, "New user has been registered.")

	//Login
	loginCredentials := []byte(`{"email_address":"kamran123@gmail.com","password":"1234444"}`)
	loginUrl := "http://localhost:8080/api/v1/user/login"
	boomRes, err = SendRequest("POST", loginUrl, loginCredentials, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(boomRes.Message)
	assert.Equal(t, boomRes.Message, "Unauthorized")
}

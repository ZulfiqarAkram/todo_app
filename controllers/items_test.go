package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"strings"
	"testing"
	"todo_app/types"
)

func TestAddItem(t *testing.T) {
	//Register
	boomRes, err := RegisterUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New user has been registered.")

	//Login Request
	boomRes, err = LoginUser()
	token := boomRes.Message
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, len(strings.Split(token, ".")), 3)

	//Add item request
	boomRes, err = AddTodo(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New todo item has been added.")
}
func TestDisplayItems(t *testing.T) {
	//Register
	boomRes, err := RegisterUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New user has been registered.")

	//Login Request
	boomRes, err = LoginUser()
	token := boomRes.Message
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, len(strings.Split(token, ".")), 3)

	//Add item request
	boomRes, err = AddTodo(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New todo item has been added.")

	//Display
	displayItemUrl := "http://localhost:8080/api/todo"
	req, err := http.NewRequest("GET", displayItemUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
	var c http.Client
	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, err, nil)
	var myTodo []types.TodoItem
	err = json.NewDecoder(res.Body).Decode(&myTodo)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, len(myTodo), 1)

}
func TestUpdateItem(t *testing.T) {
	//Register
	boomRes, err := RegisterUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New user has been registered.")

	//Login Request
	boomRes, err = LoginUser()
	token := boomRes.Message
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, len(strings.Split(token, ".")), 3)

	//Add item request
	boomRes, err = AddTodo(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New todo item has been added.")

	//Update
	updateItemUrl := "http://localhost:8080/api/todo/5"
	var updateTodoItem = []byte(`{"Text":"Say No"}`)
	req, err := http.NewRequest("PUT", updateItemUrl, bytes.NewBuffer(updateTodoItem))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
	var c http.Client
	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, err, nil)
	var myTodo types.TodoItem
	err = json.NewDecoder(res.Body).Decode(&myTodo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myTodo)
	assert.Equal(t, myTodo.Text, "Say No")

}
func TestDeleteItem(t *testing.T) {
	//Register
	boomRes, err := RegisterUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New user has been registered.")

	//Login Request
	boomRes, err = LoginUser()
	token := boomRes.Message
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, len(strings.Split(token, ".")), 3)

	//Add item request
	boomRes, err = AddTodo(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, boomRes.Message, "New todo item has been added.")

	//Delete
	addItemUrl := "http://localhost:8080/api/todo/5"
	var headersWithToken = map[string]string{
		"Content-Type": "application/json",
		"Token":        token,
	}
	boomRes, err = SendRequest("DELETE", addItemUrl, nil, headersWithToken)
	fmt.Println(boomRes.Message)
	assert.Equal(t, err, nil)
	assert.Equal(t, boomRes.Message, "Todo item has been deleted.")

}

// Helper functions
func RegisterUser() (MyBoom, error) {
	//Register Request.
	var RegisterUrl = "http://localhost:8080/api/register"
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
	fmt.Println(boomRes.Message)
	return boomRes, err
}
func LoginUser() (MyBoom, error) {
	loginCredentials := []byte(`{"email_address":"hamza123@gmail.com","password":"123456"}`)
	loginUrl := "http://localhost:8080/api/login"
	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	boomRes, err := SendRequest("POST", loginUrl, loginCredentials, headers)
	fmt.Println(boomRes.Message)
	return boomRes, err
}
func AddTodo(token string) (MyBoom, error) {
	var newTodoItem = []byte(`{"Text":"Say Hi"}`)
	addItemUrl := "http://localhost:8080/api/todo"
	var headersWithToken = map[string]string{
		"Content-Type": "application/json",
		"Token":        token,
	}
	boomRes, err := SendRequest("POST", addItemUrl, newTodoItem, headersWithToken)
	fmt.Println(boomRes.Message)
	return boomRes, err
}
func SendRequest(method string, url string, body []byte, header map[string]string) (MyBoom, error) {
	var myBoom MyBoom

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return myBoom, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	var c http.Client
	res, err := c.Do(req)
	if err != nil {
		return myBoom, err
	}
	fmt.Println(res.Body)
	err = json.NewDecoder(res.Body).Decode(&myBoom)
	return myBoom, err
}

type MyBoom struct {
	ErrorType  string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

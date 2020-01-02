package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"todo_app/auth"
	"todo_app/middleware"
)

func TestAddItem(t *testing.T) {
	// Register JWT
	auth.CreateJWTManager()
	//Middleware
	n := negroni.New()
	rr := httptest.NewRecorder()

	//Register Request
	newUser := []byte(`{"username":"moon_moon", "email_address":"moon@gmail.com","password":"123456"}`)
	regReq, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(newUser))
	if err != nil {
		fmt.Println(err.Error())
	}
	regReq.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(Register)
	n.UseHandler(middleware.NewAuthorization(handler))
	n.ServeHTTP(rr, regReq)

	var jsonRes1 boomErr
	err = json.NewDecoder(rr.Body).Decode(&jsonRes1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(jsonRes1)

	//Login Request
	validUserCredentials := []byte(`{"email_address":"moon@gmail.com","password":"123456"}`)
	loginReq, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(validUserCredentials))
	if err != nil {
		fmt.Println(err.Error())
	}
	loginReq.Header.Set("Content-Type", "application/json")
	loginHandler := http.HandlerFunc(Login)
	n.UseHandler(middleware.NewAuthorization(loginHandler))
	n.ServeHTTP(rr, loginReq)

	var jsonRes boomErr
	err = json.NewDecoder(rr.Body).Decode(&jsonRes)
	if err != nil {
		fmt.Println(err.Error())
	}

	token := jsonRes.Message

	//Add item request
	var jsonStr = []byte(`{"Text":"Say Hi"}`)
	addItemReq, err := http.NewRequest("POST", "/api/todo", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	addItemReq.Header.Set("Content-Type", "application/json")
	addItemReq.Header.Set("Token", token)
	AddHandler := http.HandlerFunc(AddItem)
	n.UseHandler(middleware.NewAuthorization(AddHandler))
	n.ServeHTTP(rr, addItemReq)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":1,"Text":"Say Hi"}`
	_, err1 := JSONBytesEqual([]byte(rr.Body.String()), []byte(expected))
	if err1 != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDisplayItems(t *testing.T) {
	token, err := LoginUser([]byte(`{"id":"1"}`))
	if err != nil {
		t.Errorf(err.Error())
	}

	req, err := http.NewRequest("GET", "/api/todo", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DisplayItems)
	//Middleware
	n := negroni.New()
	n.UseHandler(middleware.NewAuthorization(handler))
	n.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ID":1,"Text":"hello"},{"ID":2,"Text":"world"}]`
	_, err1 := JSONBytesEqual([]byte(rr.Body.String()), []byte(expected))
	if err1 != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateItem(t *testing.T) {
	var jsonStr = []byte(`{"ID":1,"Text":"No Hello"}`)

	req, err := http.NewRequest("PUT", "/api/todo/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler := http.HandlerFunc(UpdateItem)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":1,"Text":"No Hello"}`
	_, err1 := JSONBytesEqual([]byte(rr.Body.String()), []byte(expected))
	if err1 != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}

func TestDeleteItem(t *testing.T) {
	var jsonStr = []byte(`{"ID":1,"Text":"hello"}`)

	req, err := http.NewRequest("PUT", "/api/todo/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler := http.HandlerFunc(DeleteItem)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":2,"Text":"world"}`
	_, err1 := JSONBytesEqual([]byte(rr.Body.String()), []byte(expected))
	if err1 != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}

// compare two bytes json
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

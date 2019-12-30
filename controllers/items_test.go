package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDisplayItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/todo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DisplayItems)
	handler.ServeHTTP(rr, req)

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

func TestAddItem(t *testing.T) {
	var jsonStr = []byte(`{"ID":3,"Text":"Say Hi"}`)

	req, err := http.NewRequest("POST", "/api/todo", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddItem)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":3,"Text":"Say Hi"}`
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

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	d "todo_app/data"
	hf "todo_app/hp_func"
	"todo_app/types"
)

const secretKey = "my_super_secret_key"
const duration = 60 * time.Second

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Login func")
	w.Header().Set("Content-Type", "application/json")
	var user types.User
	json.NewDecoder(r.Body).Decode(&user)
	if hf.IsAuthenticUser(user.EmailAddress, user.Password) {
		fmt.Println("duration: ", duration)
		jwtManager := NewJWTWithConf(secretKey, duration)
		userInDB := hf.GetUser(user.EmailAddress, user.Password)
		payLoadData := map[string]interface{}{
			"id":            userInDB.ID,
			"username":      userInDB.Username,
			"password":      userInDB.Password,
			"email_address": userInDB.EmailAddress,
		}
		token, err := jwtManager.Sign(payLoadData)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(token)
	} else {
		json.NewEncoder(w).Encode("{'msg':'un authorized user.'}")
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser types.User
	json.NewDecoder(r.Body).Decode(&newUser)

	if newUser.EmailAddress == "" || newUser.Username == "" || newUser.Password == "" {
		json.NewEncoder(w).Encode("Please provide valid username/email_address/password.")
	} else if hf.IsDuplicateUser(newUser.EmailAddress) {
		json.NewEncoder(w).Encode("This email already exists in the system.")
	} else {
		newUser.ID = len(d.UserDB) + 1
		d.UserDB = append(d.UserDB, newUser)
		fmt.Println("New user added: ", d.UserDB)
		json.NewEncoder(w).Encode("New user has been registered.")
	}
}

func isValidToken(token string) (map[string]interface{}, error, bool) {
	jwtManager := NewJWTWithConf(secretKey, duration)
	payLoadResult, err := jwtManager.Decode(token)
	if err != nil {
		fmt.Println(err)
		return payLoadResult, err, false
	}
	fmt.Println(payLoadResult)
	return payLoadResult, err, true
}

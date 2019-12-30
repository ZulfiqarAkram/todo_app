package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"todo_app/auth"
	d "todo_app/db"
	hf "todo_app/hp_func"
	"todo_app/types"
)

//TODO shift to config
const secretKey = "my_super_secret_key"
const duration = 60 * time.Second

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Login func")
	w.Header().Set("Content-Type", "application/json")
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	if hf.IsAuthenticUser(user.EmailAddress, user.Password) {
		fmt.Println("duration: ", duration)
		jwtManager := auth.NewJWTWithConf(secretKey, duration)
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
		err1 := json.NewEncoder(w).Encode(token)
		if err1 != nil {
			fmt.Println(err1)
		}
	} else {
		err := json.NewEncoder(w).Encode("{'msg':'un authorized user.'}")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
	}

	if newUser.EmailAddress == "" || newUser.Username == "" || newUser.Password == "" {
		err := json.NewEncoder(w).Encode("Please provide valid username/email_address/password.")
		if err != nil {
			fmt.Println(err)
		}
	} else if hf.IsDuplicateUser(newUser.EmailAddress) {
		err := json.NewEncoder(w).Encode("This email already exists in the system.")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		newUser.ID = len(d.UserDB) + 1
		d.UserDB = append(d.UserDB, newUser)
		fmt.Println("New user added: ", d.UserDB)
		err := json.NewEncoder(w).Encode("New user has been registered.")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func isValidToken(token string) (map[string]interface{}, error, bool) {
	jwtManager := auth.NewJWTWithConf(secretKey, duration)
	payLoadResult, err := jwtManager.Decode(token)
	if err != nil {
		fmt.Println(err)
		return payLoadResult, err, false
	}
	fmt.Println(payLoadResult)
	return payLoadResult, err, true
}

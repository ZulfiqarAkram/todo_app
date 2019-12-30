package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo_app/auth"
	st "todo_app/store"
	"todo_app/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Login func")
	w.Header().Set("Content-Type", "application/json")
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	if IsAuthenticUser(user.EmailAddress, user.Password) {

		var usr st.MyUser
		userInDB := usr.GetUserByEmailAndPassword(user.EmailAddress, user.Password)
		payLoadData := map[string]interface{}{
			"id":            userInDB.ID,
			"username":      userInDB.Username,
			"password":      userInDB.Password,
			"email_address": userInDB.EmailAddress,
		}
		token, err := auth.JWTManager.Sign(payLoadData)
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
	} else if IsDuplicateUser(newUser.EmailAddress) {
		err := json.NewEncoder(w).Encode("This email already exists in the system.")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		st.AddUser(newUser)
		err := json.NewEncoder(w).Encode("New user has been registered.")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func IsValidToken(token string) (map[string]interface{}, error, bool) {
	payLoadResult, err := auth.JWTManager.Decode(token)
	if err != nil {
		fmt.Println(err)
		return payLoadResult, err, false
	}
	fmt.Println(payLoadResult)
	return payLoadResult, err, true
}

func IsDuplicateUser(emailAddress string) bool {
	usrInDB := st.GetUserByEmailAddress(emailAddress)
	if usrInDB.ID > 0 {
		return true
	}
	return false
}

func IsAuthenticUser(emailAddress string, password string) bool {
	usrInDB := st.GetUserByEmailAddress(emailAddress)
	return usrInDB.EmailAddress == emailAddress && usrInDB.Password == password
}

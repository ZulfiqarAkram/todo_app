package controllers

import (
	"encoding/json"
	"net/http"
	"todo_app/auth"
	st "todo_app/store"
	"todo_app/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if IsAuthenticUser(user.EmailAddress, user.Password) {
		userInDB := st.GetUserByEmailAndPassword(user.EmailAddress, user.Password)
		payLoadData := map[string]interface{}{
			"id":            userInDB.ID,
			"username":      userInDB.Username,
			"password":      userInDB.Password,
			"email_address": userInDB.EmailAddress,
		}
		token, err := auth.JWTManager.Sign(payLoadData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err1 := json.NewEncoder(w).Encode(token)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if newUser.EmailAddress == "" || newUser.Username == "" || newUser.Password == "" {
		err := json.NewEncoder(w).Encode("Please provide valid username/email_address/password.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if IsDuplicateUser(newUser.EmailAddress) {
		err := json.NewEncoder(w).Encode("This email already exists in the system.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		st.AddUser(newUser)
		err := json.NewEncoder(w).Encode("New user has been registered.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetPayload(token string) (map[string]interface{}, error) {
	payLoadResult, err := auth.JWTManager.Decode(token)
	return payLoadResult, err
}

func IsDuplicateUser(emailAddress string) bool {
	usrInDB := st.GetUserByEmailAddress(emailAddress)
	return usrInDB.ID > 0
}

func IsAuthenticUser(emailAddress string, password string) bool {
	usrInDB := st.GetUserByEmailAddress(emailAddress)
	return usrInDB.EmailAddress == emailAddress && usrInDB.Password == password
}

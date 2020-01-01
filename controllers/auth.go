package controllers

import (
	"encoding/json"
	"github.com/darahayes/go-boom"
	"net/http"
	"todo_app/auth"
	st "todo_app/store"
	"todo_app/types"
	"todo_app/validator"
)

var myValidator = validator.New()

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginUser types.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		boom.BadRequest(w, err.Error())
		//		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := myValidator.ValidateStruct(loginUser); err != nil {
		boom.BadRequest(w, err.Error())
		//		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !IsAuthenticUser(loginUser.EmailAddress, loginUser.Password) {
		boom.Unathorized(w, "Unauthorized")
		//		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userInDB := st.GetUserByEmailAndPassword(loginUser.EmailAddress, loginUser.Password)
	payLoadData := map[string]interface{}{
		"id":            userInDB.ID,
		"username":      userInDB.Username,
		"password":      userInDB.Password,
		"email_address": userInDB.EmailAddress,
	}
	token, err := auth.JWTManager.Sign(payLoadData)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//err = json.NewEncoder(w).Encode(token)
	err = JsonResponse(w, 200, token)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var newUser types.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		boom.BadRequest(w, err.Error())

		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := myValidator.ValidateStruct(newUser); err != nil {
		boom.BadRequest(w, err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if IsDuplicateUser(newUser.EmailAddress) {
		err := json.NewEncoder(w).Encode("This email already exists in the system.")
		if err != nil {
			boom.Internal(w, err.Error())
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	st.AddUser(types.User{
		ID:           0,
		Username:     newUser.Username,
		Password:     newUser.Password,
		EmailAddress: newUser.EmailAddress,
	})
	err = JsonResponse(w, 200, "New user has been registered")
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IsDuplicateUser(emailAddress string) bool {
	usrInDB := st.GetUserByEmailAddress(emailAddress)
	return usrInDB.ID > 0
}

func IsAuthenticUser(emailAddress string, password string) bool {
	usrInDB := st.GetUserByEmailAddress(emailAddress)
	return usrInDB.EmailAddress == emailAddress && usrInDB.Password == password
}

type boomErr struct {
	ErrorType  string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, message string) error {
	var codes = map[int]string{
		100: "Continue",
		101: "Switching Protocols",
		102: "Processing",
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non-Authoritative Information",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		207: "Multi-Status",
		300: "Multiple Choices",
		301: "Moved Permanently",
		302: "Moved Temporarily",
		303: "See Other",
		304: "Not Modified",
		305: "Use Proxy",
		307: "Temporary Redirect",
		400: "Bad Request",
		401: "Unauthorized",
		402: "Payment Required",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		406: "Not Acceptable",
		407: "Proxy Authentication Required",
		408: "Request Time-out",
		409: "Conflict",
		410: "Gone",
		411: "Length Required",
		412: "Precondition Failed",
		413: "Request Entity Too Large",
		414: "Request-URI Too Large",
		415: "Unsupported Media Type",
		416: "Requested Range Not Satisfiable",
		417: "Expectation Failed",
		418: "I'm a teapot",
		422: "Unprocessable Entity",
		423: "Locked",
		424: "Failed Dependency",
		425: "Unordered Collection",
		426: "Upgrade Required",
		428: "Precondition Required",
		429: "Too Many Requests",
		431: "Request Header Fields Too Large",
		451: "Unavailable For Legal Reasons",
		500: "Internal Server Error",
		501: "Not Implemented",
		502: "Bad Gateway",
		503: "Service Unavailable",
		504: "Gateway Time-out",
		505: "HTTP Version Not Supported",
		506: "Variant Also Negotiates",
		507: "Insufficient Storage",
		509: "Bandwidth Limit Exceeded",
		510: "Not Extended",
		511: "Network Authentication Required",
	}
	errorType := codes[statusCode]
	errString, _ := json.Marshal(boomErr{
		errorType,
		message,
		statusCode,
	})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(errString)
	return err
}

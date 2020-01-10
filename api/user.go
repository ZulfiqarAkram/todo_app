package api

import (
	"encoding/json"
	"github.com/darahayes/go-boom"
	"github.com/jinzhu/gorm"
	"net/http"
	"todo_app/model"
	"todo_app/util"
)

func (api *API) InitUser() {
	api.Router.User.HandleFunc("/login", api.Login).Methods(http.MethodPost)
	api.Router.User.HandleFunc("/register", api.Register).Methods(http.MethodPost)
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginUser model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		boom.BadRequest(w, err.Error())
		return
	}

	if err := api.ValidatorManager.ValidateStruct(loginUser); err != nil {
		boom.BadRequest(w, err.Error())
		return
	}

	//Is Authentic User
	usrInDB, err := api.Store.GetUserByEmailAddress(loginUser.EmailAddress)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if usrInDB.EmailAddress != loginUser.EmailAddress && usrInDB.Password != loginUser.Password {
		boom.Unathorized(w, "Unauthorized")
		return
	}

	userInDB, err := api.Store.GetUserByEmailAndPassword(loginUser.EmailAddress, loginUser.Password)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	payloadData := map[string]interface{}{
		"id":            userInDB.ID,
		"username":      userInDB.Username,
		"password":      userInDB.Password,
		"email_address": userInDB.EmailAddress,
	}
	token, err := api.JWTManager.Sign(payloadData)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	err = util.JsonResponse(w, 200, token)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser model.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		boom.BadRequest(w, err.Error())
		return
	}

	if err := api.ValidatorManager.ValidateStruct(newUser); err != nil {
		boom.BadRequest(w, err.Error())
		return
	}

	//Is Duplicate User
	usrInDB, err := api.Store.GetUserByEmailAddress(newUser.EmailAddress)
	if !gorm.IsRecordNotFoundError(err) {
		boom.BadRequest(w, err.Error())
		return
	}
	if usrInDB != nil {
		boom.BadRequest(w, "This email already exists in the system.")
		return
	}
	err = api.Store.AddUser(model.User{
		Username:     newUser.Username,
		Password:     newUser.Password,
		EmailAddress: newUser.EmailAddress,
	})
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	err = util.JsonResponse(w, 200, "New user has been registered.")
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}

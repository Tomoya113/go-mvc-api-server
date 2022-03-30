// リクエストを変換し、modelのロジックを呼び出し、viewに変換する
package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Tomoya113/go-mvc-api-server/pkg/model"
	"github.com/Tomoya113/go-mvc-api-server/pkg/view"
)

type IUserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	view  view.IUserView
	model model.IUserModel
}

func NewUserController(view view.IUserView, model model.IUserModel) UserController {
	controller := UserController{
		view:  view,
		model: model,
	}

	return controller
}

func (c UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	users, err := c.model.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := c.view.ConvertUsersToJson(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		params model.CreateUserParams
		err    error
	)

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.model.CreateUser(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := c.model.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := c.view.ConvertUserToJson(user)
	w.Write([]byte("A user is successfully created: \n"))
	w.Write(json)
}

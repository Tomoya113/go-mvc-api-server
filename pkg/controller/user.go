// リクエストを変換し、modelのロジックを呼び出し、viewに変換する
package controller

import (
	"encoding/json"
	"net/http"

	"go-mvc-api-server/pkg/model"
	"go-mvc-api-server/pkg/view"
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
	users := c.model.GetUsers()
	json, err := c.view.ConvertUsersToJson(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var params model.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("User is successfully created"))
	c.model.CreateUser(params)
}

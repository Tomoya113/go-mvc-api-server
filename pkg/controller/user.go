// リクエストを変換し、modelのロジックを呼び出し、viewに変換する
package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Tomoya113/go-mvc-api-server/pkg/model"
	"github.com/Tomoya113/go-mvc-api-server/pkg/view"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	users, err := model.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := view.ConvertUsersToJson(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		params model.CreateUserParams
		err    error
	)

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := model.CreateUser(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := model.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := view.ConvertUserToJson(user)
	w.Write([]byte("A user is successfully created: \n"))
	w.Write(json)
}

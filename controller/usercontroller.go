package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ffeliandra/eperdabe/helper"
	"github.com/ffeliandra/eperdabe/model"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	total := model.DB.Find(&users).RowsAffected
	if err := model.DB.Find(&users).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respon := &model.AllUsersWithTotal{
		Users: users,
		Total: total,
	}
	helper.ResponseJSON(w, http.StatusOK, respon)
}

func AddUser(write http.ResponseWriter, read *http.Request) {
	userInput := map[string]string{
		"nama":     "",
		"username": "",
		"password": "",
	}
	decoder := json.NewDecoder(read.Body)
	if err := decoder.Decode(&userInput); err != nil {
		helper.ResponseError(write, http.StatusBadRequest, err.Error())
		return
	}
	defer read.Body.Close()
	// hash field password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput["password"]), bcrypt.DefaultCost)
	// save to db
	user := model.User{
		Nama:     userInput["nama"],
		Username: userInput["username"],
		Password: string(hashPassword),
	}
	if err := model.DB.Create(&user).Error; err != nil {
		helper.ResponseError(write, http.StatusInternalServerError, err.Error())
		return
	}
	response := map[string]string{"status": "Sukses"}
	helper.ResponseJSON(write, http.StatusOK, response)
}

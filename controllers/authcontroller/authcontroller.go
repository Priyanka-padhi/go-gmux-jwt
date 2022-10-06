package authcontroller

import (
	"encoding/json"
	"go-gmux-jwt/helper"
	"go-gmux-jwt/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//input json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//      data user      username
	//var user models.User
	//if err := models.DB.Where("username = ?",userInput.Username).First(&user).Error;
	//

}
func Register(w http.ResponseWriter, r *http.Request) {
	//input json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//hash pass bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	//insert in db
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
func Logout(w http.ResponseWriter, r *http.Request) {
	//var userInput models.User
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&userInput); err := nil{
	//   res := map[string]string("message":err.Error())

}

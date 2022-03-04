package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:3BNuVw@1@tcp(127.0.0.1:3306)/GoUsers?parseTime=true"

type User struct {
	gorm.Model
	UserName    string `json:"UserName" binding:"required"`
	DateOfBirth string `json:"DateOfBirth" binding:"required"`
	PhoneNumber int    `json:"PhoneNumber" binding:"required"`
	EmailId     string `json:"EmailId" binding:"required"`
}

// Method to initialize the database and to enable the auto migration
func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println((err.Error()))
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&User{})
}

// Handler methods
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type set to application/json
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type set to application/json
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type set to application/json
	var user User                                      // varible user of type User (struct)
	json.NewDecoder(r.Body).Decode(&user)              // json.NewDecoder will take the input from request body and Decode to reference of user
	DB.Create(&user)                                   // user is saved to database
	json.NewEncoder(w).Encode(user)                    // user info back to browser
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type set to application/json
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type set to application/json
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is Deleted Successfully")

}

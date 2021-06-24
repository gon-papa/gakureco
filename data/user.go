package data

import (
	"log"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum"`
}

func (u *User) CreateUser(r *http.Request) {
	user := User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	validate_err := validator.New().Struct(user)
	if validate_err != nil {
		//後にvalidetioncheck関数からtemplateに返す
		log.Fatal(validate_err, "validationにひっかかりました")
	}

	err := Db.Create(&user)
	if err != nil {
		log.Fatal(err, "保存されませんでした")
	}
}

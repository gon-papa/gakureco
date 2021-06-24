package data

import (
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `validate:"-"`
	Name       string `validate:"required"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,alphanum"`
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
		fmt.Println(validate_err, "validationにひっかかりました")
		return
	}
	if err := Db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return
	}
}

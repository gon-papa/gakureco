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

func (u *User) ValidationCheck() (result map[string]string, err error) {
	result = make(map[string]string)
	err = validator.New().Struct(u)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Name":
					result["Name"] = "名前を入力してください"
					// fmt.Println(result["Name"])
				case "Email":
					result["Email"] = "正しいメールアドレスの形式で入力してください"
					// fmt.Println(result["Email"])
				case "Password":
					result["Password"] = "パスワードを入力してください"
					// fmt.Println(result["Password"])
				}
			}
		}
		return result, err
	}
	return result, err
}

func (u *User) CreateUser(r *http.Request) {
	user := &User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	s, err := user.ValidationCheck()
	if err != nil {
		fmt.Println(s, err)
		return
	}

	if err := Db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return
	}
}

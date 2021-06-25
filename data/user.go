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
	Email      string `validate:"required,email,emailunique"`
	Password   string `validate:"required,alphanum"`
}

// バリデーションチェック
func (u *User) ValidationCheck() (result map[string]string, err error) {
	result = make(map[string]string)
	validate := validator.New()
	validate.RegisterValidation("emailunique", u.EmailUnique)
	// err = validator.New().Struct(u)
	err = validate.Struct(u)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Name":
					result["Name"] = "名前を入力してください"
					// fmt.Println(result["Name"])
				case "Email":
					switch errors[i].Tag() {
					case "required":
						result["Email"] = "正しいメールアドレスの形式で入力してください"
					case "emailunique":
						result["uEmail"] = "メールアドレスが重複しています"
					}
					// result["Email"] = "正しいメールアドレスの形式で入力してください"
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

// カスタムバリデーション
func (u *User) EmailUnique(fl validator.FieldLevel) bool {
	err := Db.Where("email = ?", u.Email).Find(u)
	fmt.Println(err, "カスタムバリデーション")
	return err == nil
}

// ユーザー作成~DB保存
func (u *User) CreateUser(r *http.Request) (map[string]string, error) {
	user := &User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	var result map[string]string
	var err error

	result, err = user.ValidationCheck()
	if err != nil {
		fmt.Println(result, err)
		return result, err
	}

	if err := Db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, err
}

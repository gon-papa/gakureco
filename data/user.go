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
	Password   string `validate:"required,alphanum,min=4,max=12"`
}

// バリデーションチェック
func (u *User) ValidationCheck() (result map[string]string, err error) {
	result = make(map[string]string)
	validate := validator.New()

	validate.RegisterValidation("emailunique", u.EmailUnique)
	err = validate.Struct(u)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Name":
					result["Name"] = "名前を入力してください"
				case "Email":
					switch errors[i].Tag() {
					case "required":
						result["Email"] = "正しいメールアドレスの形式で入力してください"
					case "emailunique":
						result["uEmail"] = "メールアドレスが重複しています"
					}
				case "Password":
					result["Password"] = "パスワードを入力してください(4文字以上、12文字以内)"
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
	return err.RowsAffected == 0
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

	user.Password = Encrypt(user.Password)
	if err := Db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, err
}

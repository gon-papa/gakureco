package data

import (
	"errors"
	"net/http"
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func FindUser(email string, password string) (user User, err error) {
	result := Db.Where(&User{Email: email, Password: Encrypt(password)}).Find(&user)
	if result.RowsAffected == 0 {
		err = errors.New("EmailとPasswordの組み合わせが間違っているか、登録されていません")
		return user, err
	}
	return user, err

}

func Authenticate(r *http.Request) (User, error) {
	user, err := FindUser(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		return user, err
	}
	return user, err
}

// func (u *User) CurrentUser() bool {

// }

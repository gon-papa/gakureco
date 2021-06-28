package data

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Uuid   string
	Email  string
	UserID int
	User   User
}

func (u *User) CreateSession() (session Session, err error) {
	uu, err := createUUID()
	if err != nil {
		return
	}
	session = Session{
		Uuid:   uu,
		Email:  Encrypt(u.Email),
		UserID: int(u.Model.ID),
	}
	// fmt.Printf("%#v", session)
	Db.Create(&session)
	return session, err
}

//エラーハンドリングをする
func createUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		uu := u.String()
		return uu, err
	}
	uu := u.String()
	return uu, err
}

// 純粋にユーザー検索だけの関数にするか検討 return userとerrはuserが存在するかをerrで返す
func FindUser(email string, password string) (user User, err error) {
	result := Db.Where(&User{Email: email, Password: Encrypt(password)}).Find(&user)
	if result.RowsAffected == 0 {
		err = errors.New("EmailとPasswordの組み合わせが間違っているか、登録されていません")
		return user, err
	}
	return user, err

}

// ログイン
func Authenticate(w http.ResponseWriter, r *http.Request) (User, error) {
	user, err := FindUser(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		return user, err
	}
	session, err := user.CreateSession()
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return user, err
}

// func (u *User) CurrentUser() bool {

// }

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

// セッション作成~DB保存
func (u *User) CreateSession() (session *Session) {
	// uuidを作成
	uu, err := createUUID()
	if err != nil {
		fmt.Println(err)
	}
	// Session構造体へ値を格納し保存
	session = &Session{
		Uuid:   uu,
		Email:  Encrypt(u.Email),
		UserID: int(u.Model.ID),
		User:   *u,
	}
	// fmt.Printf("%#v", session)
	result := Db.Where(&Session{UserID: int(u.Model.ID), Email: Encrypt(u.Email)}).Find(&session)
	if result.RowsAffected == 0 {
		// fmt.Println("作成")
		Db.Create(&session)
	} else {
		// fmt.Println("更新")
		Db.Model(&session).UpdateColumn("Uuid", uu)
	}
	// 構造体を返す
	return session
}

//エラーハンドリングをする
// セッションのUUIDを生成
func createUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
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
	// ユーザーが存在するか確認
	user, err := FindUser(r.FormValue("email"), r.FormValue("password"))
	// ユーザーがいなければ、空のUser構造体とerrを返す
	if err != nil {
		return user, err
	}
	session := user.CreateSession()
	// 作成したSessionの値をCookieに保存
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	// cookieset
	http.SetCookie(w, &cookie)
	return user, err
}

func (s *Session) CurrentUser(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return false
	}
	Db.Where(&Session{Uuid: cookie.Value}).Find(s)
	uuid := s.Uuid
	return cookie.Value == uuid
}

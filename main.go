package main

import (
	"fmt"
	"gakureco/data"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

var user data.User

func index(w http.ResponseWriter, r *http.Request) {
	// ファイル解析
	t, err := template.ParseFiles("templates/index.html", "templates/layout.html")
	// エラー処理もまとめてやってくれる関数もあるが今回は初めてなので使わない(template.Must)
	if err != nil {
		log.Fatal("index.htmlが読み込めません")
	}
	// http.ResponseWriterに解析後の結果を返してレスポンスを生成
	t.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/signup.html", "templates/layout.html"))
	t.Execute(w, nil)
}

// フォーム解析、URLクエリ解析後、サインアップとログインに関数振り分け
func handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Formを解析できませんでした")
	}

	query := r.FormValue("hook")

	if query == "signup" {
		result, _ := user.CreateUser(r)
		signupRedirect(w, r, result)
	} else if query == "login" {
		fmt.Fprintln(w, "ログイン")
	} else {
		log.Fatalf("入力が間違っています。%vは不正な値です。", query)
	}
}

func signupRedirect(w http.ResponseWriter, r *http.Request, result map[string]string) {
	if len(result) != 0 {
		t := template.Must(template.ParseFiles("templates/signup.html", "templates/layout.html"))
		t.Execute(w, result)
		fmt.Println("Redirect")
		return
	}
	// 後にダッシュボードに遷移
	fmt.Fprintln(w, "OK")
}

func main() {
	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loding .env file!")
	}

	// データベース接続
	data.DatabaseConection()

	// ここでpublic傘下の静的ファイルを呼び出している。これがないとCSSが適応されない
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// マルチプレクサから各リクエストに対して関数の呼び出し登録
	http.HandleFunc("/", index)
	http.HandleFunc("/signup/", signup)
	http.HandleFunc("/usercreate", handleLogin)

	server := http.Server{
		Addr: os.Getenv("URL"),
	}
	fmt.Println("Server run with localhost:8000/")
	log.Fatal(server.ListenAndServe())
}

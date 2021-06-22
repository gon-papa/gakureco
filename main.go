package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

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

func userCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loding .env file!")
	}
	// ここでpublic傘下の静的ファイルを呼び出している。これがないとCSSが適応されない
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// マルチプレクサにハンドラをルートのリクエストに対してindexのt呼び出し登録
	http.HandleFunc("/", index)
	http.HandleFunc("/signup/", signup)
	http.HandleFunc("/usercreate/", userCreate)

	server := http.Server{
		Addr: os.Getenv("URL"),
	}
	fmt.Println("Server run with localhost:8000/")
	log.Fatal(server.ListenAndServe())
}

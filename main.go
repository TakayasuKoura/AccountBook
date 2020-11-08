package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// accountbook.dbというファイルでデータベース接続を行う
	db, err := sql.Open("sqlite3", "./accountbook.db")
	if err != nil {
		log.Fatal(err)
	}

	// AccountBookをNewAccountBookを使って作成
	ab := NewAccountBook(db)

	// テーブルを作成
	if err := ab.CreateTable(); err != nil {
		log.Fatal(err)
	}

	// HandlersをNewHandlersを使って作成
	hs := NewHandlers(ab)

	// ハンドラの登録
	http.HandleFunc("/", hs.ListHandler)
	http.HandleFunc("/save", hs.SaveHandler)
	http.HandleFunc("/summary", hs.SummaryHandler)

	fmt.Println("http://localhost:8080 で起動中...")

	// HTTPサーバを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// データベース接続、ドライバにはSQLiteを使う、ドライバ名：sqlite.DriverName
	// accountbook.dbというファイルでデータベース接続を行う
	db, err := sql.Open("sqlite3", "./accountbook.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	// AccountBookをNewAccountBookを使って作成
	ab := NewAccountBook(db)

	// テーブルを作成
	if err := ab.CreateTable(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

LOOP: // 以下のループにラベル「LOOP」をつける
	for {
		// モードを選択して実行する
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]終了")
		fmt.Printf(">")
		fmt.Scan(&mode)

		// モードによって処理を変える
		switch mode {
		case 1:
			var n int
			fmt.Print("何件入力しますか>")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				if err := ab.AddItem(inputItem()); err != nil {
					fmt.Fprintln(os.Stderr, "エラー:", err)
					break LOOP
				}
			}
		case 2:
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				break LOOP
			}
			showItems(items)
		case 3:
			fmt.Println("終了します")
			return
		}
	}
}

// Itemを入力し返す関数
func inputItem() *Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)
	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return &item
}

// Itemの一覧を出力する関数
func showItems(items []*Item) {
	fmt.Println("==========")

	// itemsの要素を1つずつ取り出してitemに入れて繰り返す
	for _, item := range items {
		fmt.Printf("[%04d] %s:%d円\n", item.ID, item.Category, item.Price)
	}

	fmt.Println("==========")
}

package main

import "fmt"

// Item ...
type Item struct {
	category string
	price    int
}

func main() {
	// 入力するデータの件数を入れる変数
	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	var items = make([]Item, 0, n)
	for i := 0; i < cap(items); i++ {
		items = inputItem(items)
	}

	// 表示
	showItems(items)
}

// 入力を行う関数
// 追加を行うItemのスライスを受け取る
// 新しく入力したItemをスライスに追加して返す
func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.category)
	fmt.Print("値段>")
	fmt.Scan(&item.price)

	// スライスに新しく入力したitemを追加する
	items = append(items, item)

	return items
}

// 一覧の表示を行う関数
func showItems(items []Item) {
	fmt.Println("==========")

	for i := 0; i < len(items); i++ {
		fmt.Printf("%s:%d円\n", items[i].category, items[i].price)
	}

	fmt.Println("==========")
}

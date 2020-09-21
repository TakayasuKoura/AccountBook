package main

import "fmt"

// Item ...
type Item struct {
	category string
	price    int
}

func main() {
	var item = inputItem()

	fmt.Println("==========")
	fmt.Printf("%sに%d円使いました\n", item.category, item.price)
	fmt.Println("==========")
}

// 入力を行う関数
// 入力したItemを返す
func inputItem() Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.category)
	fmt.Print("値段>")
	fmt.Scan(&item.price)

	return item
}

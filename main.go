package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Item ...
type Item struct {
	category string
	price    int
}

func main() {

	file, err := os.OpenFile("accountbook.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 入力するデータの件数を入れる変数
	var n int
	fmt.Print("何件入力しますか>")
	//fmt.Scan(&n)
	n = 1

	for i := 0; i < n; i++ {
		if err := inputItem(file); err != nil {
			log.Fatal(err)
		}
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	// 表示
	if err := showItems(); err != nil {
		log.Fatal(err)
	}
}

// 入力を行いファイルに保存する関数
// エラーが発生した場合にはそのまま返す
func inputItem(file *os.File) error {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.category)
	fmt.Print("値段>")
	fmt.Scan(&item.price)

	// ファイルに書き出す「品目 値段」のように書き出す
	line := fmt.Sprintf("%s %d\n", item.category, item.price)
	if _, err := file.WriteString(line); err != nil {
		return err
	}

	return nil
}

// 一覧の表示を行う関数
func showItems() error {
	// "accountbook.txt"という名前のファイルを読み込み用で開く
	file, err := os.Open("accountbook.txt")
	if err != nil {
		return err
	}

	fmt.Println("==========")

	scanner := bufio.NewScanner(file)
	// 1行ずつ読み込む
	for scanner.Scan() {
		// 1行分を取り出す
		line := scanner.Text()

		// 1行をスペースで分割する
		splited := strings.Split(line, " ")

		// 2つに分割できなかった場合はエラー
		if len(splited) != 2 {
			return errors.New("パースに失敗しました")
		}

		// 1つ目が品目
		category := splited[0]

		// 2つ目が値段
		price, err := strconv.Atoi(splited[1])
		if err != nil {
			return err
		}

		fmt.Printf("%s:%d円\n", category, price)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("==========")

	return nil
}

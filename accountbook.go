package main

import (
	"database/sql"
)

type Item struct {
	ID       int
	Category string
	Price    int
}

type Summary struct {
	Category string
	Count    int
	Sum      int
}

// AccountBook 家計簿の処理を行う型
type AccountBook struct {
	db *sql.DB
}

// NewAccountBook 新しいAccountBookを作成する関数
func NewAccountBook(db *sql.DB) *AccountBook {
	// AccountBookのポインタを返す
	return &AccountBook{db: db}
}

// CreateTable テーブルがなかったら作成する
func (ab *AccountBook) CreateTable() error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS items(
		id        INTEGER PRIMARY KEY,
		category  TEXT NOT NULL,
		price     INTEGER NOT NULL
	);`

	_, err := ab.db.Exec(sqlStr)
	if err != nil {
		return err
	}

	return nil
}

// AddItem データベースに新しいItemを追加する関数
func (ab *AccountBook) AddItem(item *Item) error {
	const sqlStr = `INSERT INTO items(category, price) VALUES(?, ?)`
	_, err := ab.db.Exec(sqlStr, item.Category, item.Price)
	if err != nil {
		return err
	}

	return nil
}

// GetItems 最近追加したものを最大limit件だけItemを取得する関数
// エラーが発生したら第2戻り値を返す
func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {
	const sqlStr = `SELECT * FROM items ORDER BY id DESC LIMIT ?`
	rows, err := ab.db.Query(sqlStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	var items []*Item
	// 1行ずつ取得した行をみる
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Category, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// DeleteItem 入力されたIDに該当するデータを削除する
func (ab *AccountBook) DeleteItem(id int) error {
	const sqlStr = `DELETE FROM items WHERE id = ?`
	_, err := ab.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}

// 集計結果を取得する関数
func (ab *AccountBook) GetSummaries() ([]*Summary, error) {
	// 品目ごとにグループ化して金額の合計を出す
	const sqlStr = `SELECT category, COUNT(*), SUM(price) FROM items GROUP BY category`
	rows, err := ab.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	var summaries []*Summary
	for rows.Next() {
		var s Summary
		err := rows.Scan(&s.Category, &s.Count, &s.Sum)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return summaries, nil
}

// 平均を取得する関数
func (s *Summary) Avg() float64 {
	if s.Count == 0 {
		return 0
	}
	return float64(s.Sum) / float64(s.Count)
}

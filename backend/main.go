package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type post struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

// var posts = []post{
// 	{ID: "1", Url: "https://test"},
// }

func initDB(db *sql.DB) error {
	const sql = `
        CREATE TABLE IF NOT EXISTS posts (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            url TEXT NOT NULL
        );`
	_, err := db.Exec(sql)
	return err
}

func main() {
	//ハンドラー関数定義
	getPosts := func(w http.ResponseWriter, _ *http.Request) {
		// wにHello from h1!と書き込んでる
		fmt.Fprint(w, posts)
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		// wにHello from h2!と書き込んでる
		fmt.Fprint(w, "Hello from h2!\n")
	}

	// パスとハンドラー結びつける
	http.HandleFunc("/", getPosts)
	http.HandleFunc("/h2", h2)

	// webサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

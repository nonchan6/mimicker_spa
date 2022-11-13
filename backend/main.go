package main

import (
	"fmt"
	"net/http"
)

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func main() {
	//ハンドラー関数定義
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		// wにHello from h1!と書き込んでる
		fmt.Fprint(w, "Hello from h1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		// wにHello from h2!と書き込んでる
		fmt.Fprint(w, "Hello from h2!\n")
	}

	// パスとハンドラー結びつける
	http.HandleFunc("/", h1)
	http.HandleFunc("/h2", h2)

	// webサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

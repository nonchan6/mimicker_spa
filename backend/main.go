package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	id      int
	name    string
	created time.Time
	updated time.Time
}

type Post struct {
	id  int    `json:"id"`
	url string `json:"url"`
}

// var posts = []post{
// 	{ID: "1", Url: "https://test"},
// }

// func ReadPosts(db *sql.DB) {
// 	var posts []Post
// 	rows, err := db.Query("select * from posts;")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for rows.Next() {
// 		post := Post{}
// 		err = rows.Scan(&post.id, &post.url)
// 		if err != nil {
// 			panic(err)
// 		}
// 		posts = append(posts, post)
// 	}
// 	rows.Close()

// 	fmt.Println(posts)
// }

func main() {
	// "mysql", "test1_user:password@/dbname"
	db, err := sql.Open("mysql", os.Getenv("TEST_USER")+":"+os.Getenv("PASSWORD")+"@tcp(db:13306)/"+os.Getenv("DATABASE"))
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	//ハンドラー関数定義
	getPosts := func(w http.ResponseWriter, _ *http.Request) {
		var posts []Post
		rows, err := db.Query("select * from posts;")
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			post := Post{}
			err = rows.Scan(&post.id, &post.url)
			if err != nil {
				panic(err)
			}
			posts = append(posts, post)
		}
		rows.Close()
		fmt.Fprint(w, posts)
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		// wにHello from h2!と書き込んでる
		fmt.Fprint(w, "Hello from h2!\n")
	}

	// パスとハンドラー結びつける
	http.HandleFunc("/posts", getPosts)
	http.HandleFunc("/h2", h2)

	// webサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

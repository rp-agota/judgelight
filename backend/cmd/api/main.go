package main

// HTTPサーバを構築する
// フロントを実装する
// APIを実装する

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/post", Receive_program)
	http.ListenAndServe(":8080", nil)
}

func Receive_program(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("program=", r.FormValue("program"))
}

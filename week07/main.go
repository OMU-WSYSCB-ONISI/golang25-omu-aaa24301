package main

import (
	"fmt"
	"net/http"
	"runtime"
  	"strconv"
)

func main() {
	// Week 07: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 07 課題")

	// 以下に実装してください
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/cal02", calamhandler)
	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func calamhandler(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println("errorだよ")
		}
		x, _ := strconv.Atoi(r.FormValue("x"))
		y, _ := strconv.Atoi(r.FormValue("y"))
		switch r.FormValue("cal0") {
		case "+":
			fmt.Fprintln(w, x+y)
		case "-":
			fmt.Fprintln(w, x-y)
		case "x":
			fmt.Fprintln(w, x*y)
		case "/":
			fmt.Fprintln(w, x/y)
		}
}

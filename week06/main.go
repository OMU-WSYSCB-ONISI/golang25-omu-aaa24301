package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	// Week 06: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 06 課題")

	// 以下に実装してください

	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/bmi", bmihandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func bmihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)

	height_m := height / 100.0;
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "BMIは")
	fmt.Fprintf(w, "%.2f", weight/(height_m*height_m))
}

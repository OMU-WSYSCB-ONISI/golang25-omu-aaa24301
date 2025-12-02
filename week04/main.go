package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Week 04: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 04 課題")

	// 以下に実装してください
	http.HandleFunc("/info", infohandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバーの起動に失敗")
	}
}

func infohandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	var time_result = time.Now().In(jst).Format("2006年01月02日 15:04:05")
	h := r.Header
	fmt.Fprintln(w, "現在の時刻は", time_result)
	fmt.Fprintln(w, "利用ブラウザは", h["User-Agent"])
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Week 03: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 03 課題")

	// 以下に実装してください
	http.HandleFunc("/webfortune", omikujihandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバーの起動に失敗")
	}
}

func omikujihandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	s := rand.New(rand.NewSource(seed))

	var fortune [4]string = [4]string{"大吉", "中吉", "吉", "凶"}
	result := fortune[s.Int31n(4)]

	fmt.Fprintf(w, "今の運勢は%sです", result)
}

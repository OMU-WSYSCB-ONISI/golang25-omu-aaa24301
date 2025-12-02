package main

import (
	"fmt"
	"net/http"
	"runtime"
  	"strconv"
	"strings"
)

func main() {
	// Week 08: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 08 課題")

	// 以下に実装してください
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/sum02", sumhandler)
	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}

func sumhandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	var count = 0;
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	var under10 = 0;
	var under20 = 0;
	var under30 = 0;
	var under40 = 0;
	var under50 = 0;
	var under60 = 0;
	var under70 = 0;
	var under80 = 0;
	var under90 = 0;
	var under100 = 0;
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		sum += tt
		count++;
		if (tt >= 90) {
			under100++;
		} else if (tt >= 80) {
			under90++;
		} else if (tt >= 70) {
			under80++;
		} else if (tt >= 60) {
			under70++;
		} else if (tt >= 50) {
			under60++;
		} else if (tt >= 40) {
			under50++;
		} else if (tt >= 30) {
			under40++;
		} else if(tt >= 20) {
			under30++;
		} else if(tt >= 10) {
			under20++;
		} else {
			under10++;
		}
	}
	average := sum / count
	fmt.Fprintln(w, "平均値は", average)
	fmt.Fprintln(w, "90点以上は",under100,"人です。")
	fmt.Fprintln(w, "80点以上は",under90,"人です。")
	fmt.Fprintln(w, "70点以上は",under80,"人です。")
	fmt.Fprintln(w, "60点以上は",under70,"人です。")
	fmt.Fprintln(w, "50点以上は",under60,"人です。")
	fmt.Fprintln(w, "40点以上は",under50,"人です。")
	fmt.Fprintln(w, "30点以上は",under40,"人です。")
	fmt.Fprintln(w, "20点以上は",under30,"人です。")
	fmt.Fprintln(w, "10点以上は",under20,"人です。")
	fmt.Fprintln(w, "0点以上は",under10,"人です。")
	fmt.Println(average)
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var time int = 0

func main() {
	// localhost:1323 に対するWebサーバを作成
	// 受けたパラメータを表示する
	http.HandleFunc("/", testHandler)

	// localhost:1323/html に対するWebサーバを作成
	// １秒間隔でクライアント側のカウンターをインクリンメントする
	http.HandleFunc("/html", printHTML)

	// port番号1323でWebサーバをスタートする
	// 開始が失敗した場合、Log情報を書き込みアプリケーションを終了する
	if err := http.ListenAndServe(":1323", nil); err != nil {
		log.Fatal("err from ListenAndServe!")
	}
}

func printHTML(w http.ResponseWriter, r *http.Request) {
	// sample.htmlをテンプレートとして読み込む
	tmpl := template.Must(template.ParseFiles("./sample.html"))

	// htmlへ送るための構造体を宣言する
	// param: Value int
	// Valueはカウントアップ用の変数
	data := &struct{
		Value int
	}{
		Value: time,
	}

	// HTMLに対して定義した構造体を埋め込む
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("execute err")
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// valueというパラメータをint型へ変換する
	queryValue, err := strconv.Atoi(r.URL.Query().Get("value"))

	// エラーが発生していた場合、アプリケーションを終了する
	if err != nil {
		log.Fatal("Atoi Error!!!")
	}

	// パラメータをresponseへ書き込む
	_, err = fmt.Fprintf(w, "%v", queryValue)

	// エラーが発生していた場合、アプリケーションを終了する
	if err != nil {
		log.Fatal("Fprintf Error!!!")
	}

	// 受けたパラメータを表示
	fmt.Println("param is ", queryValue)

	time = queryValue
}

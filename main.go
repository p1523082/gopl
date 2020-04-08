package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"gopkg.in/ini.v1"
)

var time int
var port string

func main() {
	// configファイル config.iniを読み込む
	conf, err := ini.Load("./config/config.ini")

	if err != nil {
		log.Println(err)
	}

	// webセクションのportを取得し、port変数に代入
	port = conf.Section("web").Key("port").String()

	// localhost:1323 に対するWebサーバを作成
	// 受けたパラメータを表示する
	http.HandleFunc("/", incrementHandler)

	// localhost:1323/html に対するWebサーバを作成
	// １秒間隔でクライアント側のカウンターをインクリンメントする
	http.HandleFunc("/html", printHTML)

	// port番号1323でWebサーバをスタートする
	// 開始が失敗した場合、Log情報を書き込みアプリケーションを終了する
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("err from ListenAndServe!")
	}
}

func printHTML(w http.ResponseWriter, r *http.Request) {

	// sample.htmlをテンプレートとして読み込む
	tmpl := template.Must(template.ParseFiles("./sample.html"))

	// htmlへ送るための構造体を宣言する
	// param1: Value int
	// param2: Port string
	// Valueはカウントアップ用の変数
	data := &struct{
		Value int
		Port string
	}{
		Value: time,
		Port: port,
	}

	// HTMLに対して定義した構造体を埋め込む
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("execute err")
	}
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	// アクセス一回につき、timeをインクリメントする
	time += 1

	fmt.Printf("now value is %d\n",time)

	// パラメータをresponseへ書き込む
	_, err := fmt.Fprintf(w, "%v", time)

	// エラーが発生していた場合、アプリケーションを終了する
	if err != nil {
		log.Fatal("Fprintf Error!!!")
	}
}

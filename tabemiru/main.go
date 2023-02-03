package main

import (
	"fmt"
	"log"
	"net/http"
	module "tabemiru/model"
)

func Handlerpopular(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// サービスを取得
	var service = module.GetService(*r)
	// DB接続情報取得

	module.LoadConfig()
	// DBオープン
	// db := connectionDB(conf)
	// サービスを実行
	service.Excute()

	// path := "/RecipeList"
	// key := "NewRecipe"
	// pos := strings.LastIndex(key, ".")
	// if pos >= 0 {
	// 	str := key[pos:]
	// 	if (len(str) > 0) && (str != ".js") && (str != ".css") && (str != ".map") {
	// 		path += "/" + key
	// 	}
	// } else {
	// 	path += "/" + key
	// }
	// fs := http.StripPrefix(path, http.FileServer(http.Dir("web")))
	// fs.ServeHTTP(w, r)
}
func main() {
	fmt.Println("main")

	module.LoadConfig()
	fmt.Println(module.GetDBconfig())
	// "user-form"へのリクエストを関数で処理する
	// レシピ情報取得時
	http.HandleFunc("/", Handlerpopular)

	// サーバーを起動 ※サーバー起動失敗時のログを出力するためlog.Fatalを使用
	log.Fatal(http.ListenAndServe(":1208", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	RecipeModel "rakushiru/src/Module/Model"
)

// 入力フォーム画面
func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// // テンプレートをパースする
	// tpl := template.Must(template.ParseFiles("templates/user-form.html"))

	// // テンプレートに出力する値をマップにセット
	// values := map[string]string{}

	// // マップを展開してテンプレートを出力する
	// if err := tpl.ExecuteTemplate(w, "user-form.html", values); err != nil {
	// 	fmt.Println(err)
	// }
	type Ping struct {
		Status int
		Rssult string
	}

	//Const.getMaterialStub()

	// ping := Ping{http.StatusOK, "ok"}

	res, err := json.Marshal(RecipeModel.GetModels())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// 入力内容の確認画面
func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	type Ping struct {
		Status int
		Rssult string
	}

	ping := Ping{http.StatusOK, "ok"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(res)
	w.Write(res)

	// // テンプレートをパースする
	// tpl := template.Must(template.ParseFiles("templates/user-confirm.html"))

	// // テンプレートに出力する値をマップにセット
	// values := map[string]string{
	// 	"account": req.FormValue("account"),
	// 	"name":    req.FormValue("name"),
	// 	"passwd":  req.FormValue("passwd"),
	// }

	// // マップを展開してテンプレートを出力する
	// if err := tpl.ExecuteTemplate(w, "user-confirm.html", values); err != nil {
	// 	fmt.Println(err)
	// }
}

func main() {
	fmt.Println("main")

	// "user-form"へのリクエストを関数で処理する
	http.HandleFunc("/user-form", HandlerUserForm)
	fmt.Println("a")

	// "user-confirm"へのリクエストを関数で処理する
	http.HandleFunc("/user-confirm", HandlerUserConfirm)
	fmt.Println("b")

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
	fmt.Println("c")
}

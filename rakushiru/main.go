package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	RecipeModel "rakushiru/src/Module/Model"
)

// 構造を宣言
type Request struct {
	reqCode string `json:"name"`
	data    int    `json:"age"`
}

type Response struct {
	Status int
	Rssult RecipeModel.Models
}

type EResponse struct {
	Status int
}

// 入力フォーム画面
func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	body, e := ioutil.ReadAll(r.Body)
	if e != nil {
		fmt.Printf("ERROR1")
		fmt.Println(e.Error())
		return
	}
	req := Request{}
	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("ERROR2")
		fmt.Println(e.Error())
		return
	}

	if req.reqCode == "RecipesInfoSelect" {
		res, err := json.Marshal(RecipeModel.GetModels())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%#v\n", req)
		fmt.Printf("%#v\n", req.reqCode)
		fmt.Println("OK")

		w.Write(res)
	} else {
		res, err := json.Marshal(EResponse{http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%#v\n", req)
		fmt.Printf("%#v\n", req.reqCode)
		fmt.Println("NG")

		w.Write(res)
	}
}

// 入力内容の確認画面
func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// type Ping struct {
	// 	Status int
	// 	Rssult string
	// }

	// ping := Ping{http.StatusOK, "ok"}

	// res, err := json.Marshal(ping)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// fmt.Println(res)
	// w.Write(res)

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

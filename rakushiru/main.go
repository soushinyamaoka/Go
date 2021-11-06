package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	Const "rakushiru/src/Module/Const"
	RecipeModel "rakushiru/src/Module/Model"
	Service "rakushiru/src/Module/Service"
)

// 構造を宣言
type Request struct {
	ReqCode string             `json:"reqCode"`
	Data    RecipeModel.Models `json:"data"`
}

type Response struct {
	Status int
	Data   RecipeModel.Models
}

type EResponse struct {
	Status int
}

// レシピ入力画面
func HandlerRecipeInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	body, e := ioutil.ReadAll(r.Body)
	fmt.Printf("%#v\n", string(body))
	if e != nil {
		fmt.Printf("ERROR1")
		fmt.Println(e.Error())
		return
	}
	req := Request{}
	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}

	fmt.Println("OK")
	model := RecipeModel.Models{}
	var res = Service.Result{}
	fmt.Println(req.Data.Recipes[0].RecipeId)
	model.Recipes = req.Data.Recipes
	fmt.Println("OK2")

	if req.ReqCode == Const.REQ_SAVE_RECIPE {
		fmt.Println("CALL:SaveRecipe")
		// レシピ保存処理
		res = Service.SaveRecipe(model)
	} else if req.ReqCode == Const.REQ_SEARCH_RECIPE {
		// レシピ検索処理
		fmt.Println("CALL:SearchRecipe")
		res = Service.SearchRecipe(&model)
	}

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println(res.Message)
	} else {
		res, err := json.Marshal(Response{0, model})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%#v\n", req)
		fmt.Printf("%#v\n", req.ReqCode)
		fmt.Printf("%#v\n", req.Data)
		fmt.Println("OK")

		w.Write(res)
	}
}

// レシピ入力画面
func HandlerRecipeSave(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	body, e := ioutil.ReadAll(r.Body)
	fmt.Printf("%#v\n", string(body))
	if e != nil {
		fmt.Printf("ERROR1")
		fmt.Println(e.Error())
		return
	}
	req := Request{}
	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}

	fmt.Println("OK")
	model := RecipeModel.Models{}
	var res = Service.Result{}
	fmt.Println(req.Data.Recipes[0].RecipeId)
	model.Recipes = req.Data.Recipes
	fmt.Println("OK2")

	fmt.Println("CALL:SaveRecipe")
	// レシピ保存処理
	res = Service.SaveRecipe(model)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println(res.Message)
	} else {
		res, err := json.Marshal(Response{0, model})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}

// レシピ入力画面
func HandlerRecipeMake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	body, e := ioutil.ReadAll(r.Body)
	fmt.Printf("%#v\n", string(body))
	if e != nil {
		fmt.Printf("ERROR1")
		fmt.Println(e.Error())
		return
	}
	req := Request{}
	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}

	var res = Service.Result{}
	var model = RecipeModel.Models{}
	fmt.Println("CALL:SaveRecipe")
	// レシピ作成前処理
	res = Service.MakeRecipe(&model)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println(res.Message)
	} else {
		res, err := json.Marshal(Response{0, model})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}

func main() {
	fmt.Println("main")

	// "user-form"へのリクエストを関数で処理する
	http.HandleFunc("/recipeInfo", HandlerRecipeInfo)
	http.HandleFunc("/saveRecipe", HandlerRecipeSave)
	http.HandleFunc("/makeRecipe", HandlerRecipeMake)
	// model := RecipeModel.Models{}
	// Service.SearchRecipe(&model)
	fmt.Println("a")
	// fmt.Println(model)

	fmt.Println("b")

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
	fmt.Println("c")
}

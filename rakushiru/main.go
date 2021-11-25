package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	Const "rakushiru/src/Module/Const"
	RecipeModel "rakushiru/src/Module/Model"
	Module "rakushiru/src/Module/Service"
	"strings"
)

// 構造を宣言
type Request struct {
	ReqCode string             `json:"reqCode"`
	Data    RecipeModel.Models `json:"data"`
}

type HomeRequest struct {
	ReqCode string             `json:"reqCode"`
	Data    []RecipeModel.Data `json:"data"`
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
	var res = Module.Result{}

	fmt.Println(req.Data.Recipes[0].RecipeId)
	model = req.Data
	fmt.Println("OK2")

	// レシピ検索処理
	fmt.Println(model)
	res, model = Module.OpenRecipeInfo(model)

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
	fmt.Println("call: HandlerRecipeSave")
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	req := Request{}
	fmt.Println(r)
	fmt.Println(req)

	var recipeId = ""
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("ファイルの確保失敗")
	}

	body := r.FormValue("Data")
	json.Unmarshal([]byte(body), &req)

	model := req.Data
	var res = Module.Result{}
	model.Recipes = req.Data.Recipes

	fmt.Println("CALL:SaveRecipe")
	// レシピ保存処理
	res, recipeId = Module.SaveRecipe(model)

	//アップロードされたファイル名を取得
	fmt.Println("CALL:FileUpload")
	fmt.Println(file)
	Module.FileUpload(file, recipeId)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println("エラーが発生しました")
		fmt.Println(res.Message)
	} else {
		fmt.Println("正常終了")
		res, err := json.Marshal(Response{0, model})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}

// レシピ入力画面
func HandlerOpenHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call: HandlerOpenHome")
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
	req := HomeRequest{}
	fmt.Println("body")
	fmt.Println(body)

	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}

	var model = RecipeModel.Models{}
	var res = Module.Result{}
	keyModel := req.Data

	// レシピ保存処理
	// var sample []string
	// sample = append(sample, "ラーメン")
	// sample = append(sample, "うどん")
	// string a := ["a", "b"]
	fmt.Println(keyModel)
	res, model = Module.SearchRecipe(keyModel)
	// fmt.Println(a[0])
	// model, res = Module.SearchHome(model)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println("エラーが発生しました")
		fmt.Println(res.Message)
	} else {
		fmt.Println("正常終了")
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
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// body, e := ioutil.ReadAll(r.Body)
	// fmt.Printf("%#v\n", string(body))
	// if e != nil {
	// 	fmt.Printf("ERROR1")
	// 	fmt.Println(e.Error())
	// 	return
	// }
	req := Request{}
	body := r.FormValue("Data")
	json.Unmarshal([]byte(body), &req)

	// e = json.Unmarshal(body, &req)
	// if e != nil {
	// 	fmt.Printf("%#v\n", "ERROR2")
	// 	fmt.Printf("%#v\n", e.Error())
	// 	fmt.Printf("%#v\n", string(body))
	// 	return
	// }

	var res = Module.Result{}
	var model = RecipeModel.Models{}
	fmt.Println("CALL:SaveRecipe")
	// レシピ作成前処理
	model, res = Module.MakeRecipe(model)

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

const (
	defaultMaxMemory = 3200 << 20 // 3200 MB
)

// 画像保存
func HandlerOpenImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("saveImage")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	var Path = ""

	sub := strings.TrimPrefix(r.URL.Path, "/image")
	_, id := filepath.Split(sub)
	if id != "" {
		Path += "image/" + id
	}

	img, err := os.Open(Path)
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)

	// res, err := json.Marshal(Response{0, RecipeModel.Models{}})
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(res)
}

func main() {
	fmt.Println("main")

	// "user-form"へのリクエストを関数で処理する
	http.HandleFunc("/recipeInfo", HandlerRecipeInfo)
	http.HandleFunc("/saveRecipe", HandlerRecipeSave)
	http.HandleFunc("/makeRecipe", HandlerRecipeMake)
	http.HandleFunc("/openHome", HandlerOpenHome)
	http.HandleFunc("/image/", HandlerOpenImage)
	fmt.Println("a")

	fmt.Println("b")

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
	fmt.Println("c")
}

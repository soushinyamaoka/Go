package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

type HomeResponse struct {
	Status int
	Data   RecipeModel.HomeModels
}

type EResponse struct {
	Status int
}

// レシピ入力画面
func HandlerRecipeInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerRecipeInfo")
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
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}
	model := RecipeModel.Models{}
	var res = Module.Result{}
	model = req.Data

	// レシピ検索処理
	res, model = Module.OpenRecipeInfo(model)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Printf("%#v\n", string(body))
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

	// レシピ保存処理
	res, recipeId = Module.SaveRecipe(model)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println("エラーが発生しました")
		fmt.Println(res.Message)
		res, err := json.Marshal(Response{res.Status, model})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}

	//アップロードされたファイル名を取得
	if file != nil {
		Module.FileUpload(file, recipeId)
	}

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
		fmt.Println(res)
		w.Write(res)
	}
}

// レシピ入力画面
func HandlerSearchRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call: HandlerSearchRecipe")
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

	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}

	var model = RecipeModel.Models{}
	var res = Module.Result{}

	if req.ReqCode == Const.REQ_SEARCH_NEW_RECIPE {
		// 新着レシピの場合
		res, model = Module.SearchNewRecipe()

	} else {
		// キーワード検索の場合
		keyModel := req.Data
		res, model = Module.SearchRecipe(keyModel, true)
	}

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println("エラーが発生しました")
		fmt.Println("body")
		fmt.Println(body)
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
func HandlerSearchHomeRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call: HandlerSearchRecipe")
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

	e = json.Unmarshal(body, &req)
	if e != nil {
		fmt.Printf("%#v\n", "ERROR2")
		fmt.Printf("%#v\n", e.Error())
		fmt.Printf("%#v\n", string(body))
		return
	}

	var model = RecipeModel.HomeModels{}
	var res = Module.Result{}

	keyModel := req.Data
	res, model = Module.SearchHomeRecipe(keyModel, true)

	if res.Status != Const.STATUS_SUCCESS {
		fmt.Println("エラーが発生しました")
		fmt.Println(res.Message)
		fmt.Println("body")
		fmt.Println(body)
	} else {
		fmt.Println("正常終了")
		res, err := json.Marshal(HomeResponse{0, model})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}

// レシピ入力画面
func HandlerRecipeMake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerRecipeMake")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	req := Request{}
	body := r.FormValue("Data")
	json.Unmarshal([]byte(body), &req)

	var res = Module.Result{}
	var model = RecipeModel.Models{}
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
	fmt.Println("HandlerOpenImage")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	var Path = ""

	sub := strings.TrimPrefix(r.URL.Path, "/image")
	_, id := filepath.Split(sub)
	if id != "" {
		Path += "image/" + id
	} else {
		Path += "image/noImage.jpg"
	}
	img, err := os.Open(Path)
	if err != nil {
		fmt.Println(err)
		img, err = os.Open("image/noImage.jpg")
		if err != nil {
			fmt.Println(err)
		}
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)

}

func main() {
	fmt.Println("main")

	// "user-form"へのリクエストを関数で処理する
	// レシピ情報取得時
	http.HandleFunc("/recipeInfo", HandlerRecipeInfo)
	// レシピ保存時
	http.HandleFunc("/saveRecipe", HandlerRecipeSave)
	// http.HandleFunc("/makeRecipe", HandlerRecipeMake)
	// ホーム画面表示時
	http.HandleFunc("/openHome", HandlerSearchHomeRecipe)
	// レシピ検索時
	http.HandleFunc("/searchRecipe", HandlerSearchRecipe)
	// 新着レシピ検索
	http.HandleFunc("/searchNewRecipe", HandlerSearchRecipe)
	// 画像取得
	http.HandleFunc("/image/", HandlerOpenImage)
	fmt.Println("a")

	fmt.Println("b")

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
	fmt.Println("c")
}

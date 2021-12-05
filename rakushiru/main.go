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
	"text/template"
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
		fmt.Println(res.Err)
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
		fmt.Println(res.Err)
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
		fmt.Println(res.Err)
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
		fmt.Println(res.Err)
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
	fmt.Println("call: HandlerSearchHomeRecipe")
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
		fmt.Println(res.Err)
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
		fmt.Println(res.Err)
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

// 画像保存
func HandlerNoRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerNoRoot")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	sub := strings.TrimPrefix(r.URL.Path, "/RecipeList")
	_, key := filepath.Split(sub)
	fmt.Println("key")
	fmt.Println(key)

	params := map[string]string{
		"KeyWord": "ラーメン",
	}
	tpl := template.Must(template.ParseFiles("web/request.html"))
	fmt.Println("Executeします")
	err := tpl.Execute(w, params) // dataは渡さない
	if err != nil {
		// log.Fatalln(err)
	}

	// fmt.Println("ここから")
	// fs := http.FileServer(http.Dir("web"))
	// fs := http.StripPrefix("/RecipeList/", http.FileServer(http.Dir("web")))
	// var res = Module.Result{}
	// // 新着レシピの場合
	// res, _ = Module.SearchNewRecipe()
	// res.Message = "OKです"
	// res.Status = Const.STATUS_SUCCESS
	// var store = sessions.NewCookieStore([]byte("random-string"))
	// session, _ := store.Get(r, "goapp-p-register")
	// session.Values["keyWord"] = "ラーメン"
	// session.Save(r, w)
	// fmt.Println("セッションに保存")
	// params := map[string]string{
	// 	"Name": "たろう",
	// 	"Age":  "23",
	// }
	// tpl := template.Must(template.ParseFiles("web/request.html"))
	// err := tpl.Execute(w, params) // dataは渡さない
	// if err != nil {
	// 	// log.Fatalln(err)
	// }

	// if res.Status != Const.STATUS_SUCCESS {
	// 	fmt.Println(res.Message)
	// 	fmt.Println(res.Err)
	// } else {
	// 	res, err := json.Marshal(res)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Write(res)
	// }

	// fmt.Println("ServeHTTPします")
	// fs.ServeHTTP(w, r)

	// t, err := template.ParseFiles("web/index.html")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// if err := t.Execute(w, nil); err != nil {
	// 	panic(err.Error())
	// }
	// var Path = ""

	// sub := strings.TrimPrefix(r.URL.Path, "/image")
	// _, id := filepath.Split(sub)
	// if id != "" {
	// 	Path += "image/" + id
	// } else {
	// 	Path += "image/noImage.jpg"
	// }
	// img, err := os.Open(Path)
	// if err != nil {
	// 	fmt.Println(err)
	// 	img, err = os.Open("image/noImage.jpg")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// defer img.Close()
	// w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	// io.Copy(w, img)

}

// 画像保存
func HandlerNoRoot2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerNoRoot2")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	sub := strings.TrimPrefix(r.URL.Path, "/RecipeList")
	_, key := filepath.Split(sub)
	fmt.Println("key")
	fmt.Println(key)

	// params := map[string]string{
	// 	"KeyWord": key,
	// }
	// tpl := template.Must(template.ParseFiles("web/request.html"))
	// fmt.Println("Executeします")
	// err := tpl.Execute(w, params) // dataは渡さない
	// if err != nil {
	// 	// log.Fatalln(err)
	// }
	// path := "/RecipeList"
	// chk:= ""
	// if ((len(key) > 0)) {
	// 	chk = key[0:4]
	// }

	// if (len(key) > 0) && (chk != "main") {
	// 	path += "/" + key
	// }
	// fmt.Println("ここから")
	// fmt.Println(path)
	// fs := http.FileServer(http.Dir("web"))
	fs := http.StripPrefix("/RecipeList", http.FileServer(http.Dir("web")))
	//fs := http.StripPrefix("/", http.FileServer(http.Dir("web")))

	// var res = Module.Result{}
	// // // 新着レシピの場合
	// // res, _ = Module.SearchNewRecipe()
	// res.Message = "OKです"
	// res.Status = Const.STATUS_SUCCESS

	fmt.Println("ServeHTTPします")
	fs.ServeHTTP(w, r)

}

// レシピ一覧取得(NotFound)
func NotFoundRecipeList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundRecipeList")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	sub := strings.TrimPrefix(r.URL.Path, "/RecipeList")
	_, key := filepath.Split(sub)
	fmt.Println("key")
	fmt.Println(key)

	path := "/RecipeList"
	pos := strings.LastIndex(key, ".")
	if pos >= 0 {
		str := key[pos:]
		if (len(str) > 0) && (str != ".js") && (str != ".css") && (str != ".map") {
			path += "/" + key
		}
	} else {
		path += "/" + key
	}
	fs := http.StripPrefix(path, http.FileServer(http.Dir("web")))
	fs.ServeHTTP(w, r)

}

// 新着レシピ取得(NotFound)
func NotFoundNewRecipeList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundNewRecipeList")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	key := "NewRecipe"
	fmt.Println("key")
	fmt.Println(key)

	path := "/RecipeList"
	pos := strings.LastIndex(key, ".")
	if pos >= 0 {
		str := key[pos:]
		if (len(str) > 0) && (str != ".js") && (str != ".css") && (str != ".map") {
			path += "/" + key
		}
	} else {
		path += "/" + key
	}
	fs := http.StripPrefix(path, http.FileServer(http.Dir("web")))
	fs.ServeHTTP(w, r)

}

func NotFoundRecipeInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundRecipeInfo")
	// w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	sub := strings.TrimPrefix(r.URL.Path, "/RecipeList")
	_, key := filepath.Split(sub)
	fmt.Println("key")
	fmt.Println(key)

	path := "/RecipesInfo"
	pos := strings.LastIndex(key, ".")
	if pos >= 0 {
		str := key[pos:]
		if (len(str) > 0) && (str != ".js") && (str != ".css") && (str != ".map") {
			path += "/" + key
		}
	} else {
		path += "/" + key
	}
	fs := http.StripPrefix(path, http.FileServer(http.Dir("web")))
	fs.ServeHTTP(w, r)

}

func main() {
	fmt.Println("main")

	// "user-form"へのリクエストを関数で処理する
	// レシピ情報取得時
	http.HandleFunc("/recipeInfo", HandlerRecipeInfo)
	// レシピ保存時
	http.HandleFunc("/saveRecipe", HandlerRecipeSave)
	// ホーム画面表示時
	http.HandleFunc("/openHome", HandlerSearchHomeRecipe)
	// レシピ検索時
	http.HandleFunc("/searchRecipe", HandlerSearchRecipe)
	// 新着レシピ検索
	http.HandleFunc("/searchNewRecipe", HandlerSearchRecipe)
	// 画像取得
	http.HandleFunc("/image/", HandlerOpenImage)
	fmt.Println("a")
	// NotFound
	http.HandleFunc("/RecipeList/NewRecipe/", NotFoundNewRecipeList)
	http.HandleFunc("/RecipeList/", NotFoundRecipeList)
	http.HandleFunc("/RecipesInfo/", NotFoundRecipeInfo)

	fmt.Println("b")

	// buildフォルダを公開
	// StripPrefixの第一引数がURL側、HandleがWEBのルート
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
	http.Handle("/MakeRecipes", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
	// http.Handle("/dummyForm", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
	// http.Handle("/RecipeList/", http.StripPrefix("/RecipeList/", http.FileServer(http.Dir("web"))))

	// サーバーを起動
	http.ListenAndServe(":1208", nil)
	fmt.Println("c")
}

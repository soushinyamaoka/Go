package main

import (
	"encoding/json"
	"net/http"
	b "tabemiru/module/base"
	c "tabemiru/module/common"
	home "tabemiru/module/home"
	search "tabemiru/module/search"
)

/*
名称：メインモジュール
概要：メインとなるモジュール
*/

// 構造体
type BaseStruct struct {
	Req    http.Request // リクエスト
	DB     *b.DB        // DB接続情報
	status int
}

// サービスのインターフェース
type Service interface {
	ServiceInit() error // サービス実行前処理
	Excute() error      // サービス実行処理
	Terminate() error   // サービス実行後処理
	GetRes() c.Response
}

// HomeReqの型定義
type HomeReq home.HomeSt

// SearchReqの型定義
type SearchReq search.SearchReq

var log *c.Log

func init() {
	c.LoadConfig()
}

/*
名称：メイン処理
概要：go実行時の処理
param : 無し
return : 無し
*/
func main() {

	// "user-form"へのリクエストを関数で処理する
	// レシピ情報取得時
	http.HandleFunc("/", Handler)

	// サーバーを起動
	log.Fatal("", http.ListenAndServe(":1208", nil))
}

/*
名称：レスポンス処理
概要：リクエストを元に処理を行い、レスポンスを返す
param : http.ResponseWriter
return : http.Request
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	log = c.NewLog()
	log.Info("call:Handler")

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// サービスを取得
	res, err := json.Marshal(HandlerMain(r))
	if err != nil {
		log.Error("システムエラーが発生しました: %v", err)
	}
	print(string(res))
	w.Write(res)

}

/*
名称：サービス取得処理
概要：リクエストURIに該当するサービスを返す
param : http.Request httpリクエスト
return : Service サービス
*/
func HandlerMain(r *http.Request) c.Response {
	// サービスを取得
	var service = GetService(*r)

	defer func() {
		if err := recover(); err != nil {
			log.Error("Panic occurred: %v", err)
		}
		// サービスを終了
		if err := service.Terminate(); err != nil {
			log.Error("システムエラーが発生しました: %v", err)
			return
		}
	}()

	// サービスを初期化
	if err := service.ServiceInit(); err != nil {
		log.Error("システムエラーが発生しました: %v", err)
	}

	// サービスを実行
	if err := service.Excute(); err != nil {
		log.Error("システムエラーが発生しました: %v", err)
	}

	return service.GetRes()
}

/*
名称：サービス取得処理
概要：リクエストURIに該当するサービスを返す
param : http.Request httpリクエスト
return : Service サービス
*/
func GetService(r http.Request) Service {
	if c.URL_PATH_HOOME == r.RequestURI {
		return &home.HomeSt{Req: r, DB: &b.DB{Con: nil}}
		// } else if c.URL_PATH_SEARCH == r.RequestURI {
		// 	return search.SearchReq{Req: r} //
	} else {
		return &home.HomeSt{Req: r, DB: &b.DB{Con: nil}}
	}
}

/*
名称：サービス実行前処理
概要：サービスの初期処理を行う
param : 無し
return : error エラー
*/
func (s *BaseStruct) ServiceInit() error {
	return nil
}

/*
名称：サービス実行処理
概要：サービス実行処理のインターフェース
param : 無し
return : 無し
*/
func (s BaseStruct) Excute() error {
	return nil
}

/*
名称：サービス実行後処理
概要：サービス実行後の処理を行う
param : 無し
return : error エラー
*/
func (s *BaseStruct) Terminate() error {
	return nil
}

func (s *BaseStruct) GetRes() c.Response {
	return c.Response{}
}

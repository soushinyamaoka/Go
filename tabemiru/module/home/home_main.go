package home

import (
	"net/http"
	b "tabemiru/module/base"
	c "tabemiru/module/common"
	e "tabemiru/module/entity"
)

/*
名称：ホーム画面用メインモジュール
概要：ホーム画面内での処理を行うメインモジュール
*/

// 構造体
type HomeSt struct {
	Req    http.Request  // リクエスト
	DB     *b.DB         // DB接続情報
	en     e.YoutubeInfo // youtube_infoテーブル取得情報
	status int
}

// サービスのインターフェース
type Service interface {
	ServiceInit() error
	Excute() error
	Terminate() error
	GetRes() c.Response
}

/*
名称：サービス初期処理
概要：ホーム画面表示時処理の初期処理を行う
param : 無し
return : error エラー
*/
func (h *HomeSt) ServiceInit() error {
	// DB接続
	if err := h.Open(); err != nil {
		// エラーステータスをセット
		h.status = c.STATUS_DB_ERROR
		return err
	}

	return nil
}

/*
名称：サービス実行処理
概要：ホーム画面表示時の処理を行う
param : 無し
return : error エラー
*/
func (h *HomeSt) Excute() error {

	// Youtube情報取得
	if err := h.SelectYoutubeInfo(); err != nil {
		// エラーステータスをセット
		h.status = c.STATUS_UNEXPECTED
		return err
	}

	// 成功ステータスをセット
	h.status = c.STATUS_SUCCESS

	return nil
}

/*
名称：サービス実行後処理
概要：ホーム画面表示時処理実行後の処理を行う
param : 無し
return : error エラー
*/
func (h *HomeSt) Terminate() error {
	// DB切断
	if err := h.Close(); err != nil {
		h.status = c.STATUS_DB_ERROR
		return err
	}

	return nil
}

func (s *HomeSt) GetRes() c.Response {
	// エラーが発生している場合
	if s.status != c.STATUS_SUCCESS {
		return c.Response{Status: s.status,
			Data: nil}
		// 成功している場合
	} else {
		return c.Response{Status: s.status,
			Data: s.en}
	}
}

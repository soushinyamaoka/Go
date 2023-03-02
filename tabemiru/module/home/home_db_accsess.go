package home

import (
	c "tabemiru/module/common"

	_ "github.com/go-sql-driver/mysql"
)

/*
名称：ホーム画面用メインモジュール
概要：ホーム画面内での処理を行うメインモジュール
*/

var err error

/*
名称：DB接続
概要：DB接続処理を呼び出す
param : 無し
return : 無し
*/
func (homeSt *HomeSt) Open() error {
	return homeSt.DB.Open()
}

/*
名称：DB切断
概要：DBの切断処理を呼び出す
param : 無し
return : 無し
*/
func (homeSt *HomeSt) Close() error {
	return homeSt.DB.Close()
}

/*
名称：Youtube情報取得処理
概要：DBからYoutube情報を取得する
param : 無し
return : 無し
*/
func (homeSt *HomeSt) SelectYoutubeInfo() error {
	c.GetLog().Debug("Call:SelectYoutubeInfo")

	// DB検索実行
	if err := homeSt.DB.Con.Debug().Find(&homeSt.en).Error; err != nil {
		return err
	}
	return nil
}

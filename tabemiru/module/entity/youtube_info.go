package base

import "time"

/*
  名称：youtube_infoテーブル構造体
  概要：youtube_infoテーブル構造体メインモジュール
*/

// 構造体
type YoutubeInfo struct {
	Recipe_id   string    `gorm:"recipe_id" json:"RecipeId,omitempty"`
	Youtube_id  string    `gorm:"youtube_id" json:"Youtube_id,omitempty"`
	Title       string    `gorm:"title" json:"Title,omitempty"`
	Del_flg     string    `gorm:"del_flg" json:"Del_flg,omitempty"`
	Regist_date time.Time `gorm:"regist_date" json:'Regist_date'`
	Update_date time.Time `gorm:"update_date" json:'Update_date'`
}

// テーブル名インターフェース
type Tabler interface {
	TableName() string
}

/*
名称：テーブル名取得処理
概要：DB接続時のテーブル名を取得する
param : 無し
return : テーブル名
*/
func (YoutubeInfo) TableName() string {
	return "youtube_info"
}

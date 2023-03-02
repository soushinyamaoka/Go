package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
名称：サービスベースモジュール
概要：各サービスモジュールのベースとなるモジュール
*/

// 構造体
type Config struct {
	ServeInfo ServeConf `json:'serveInfo'`
	DbInfo    DBConf    `json:'dbInfo'`
	SiteInfo  SiteConf  `json:'siteInfo'`
}

type ServeConf struct {
	Host            string `json:'host'`
	Port            string `json:'port'`
	RecipeInfo      string `json:'recipeInfo'`
	SearchRecipe    string `json:'searchRecipe'`
	SearchNewRecipe string `json:'searchNewRecipe'`
	SaveRecipe      string `json:'saveRecipe'`
	SaveImage       string `json:'saveImage'`
	OpenHome        string `json:'openHome'`
	OpenImage       string `json:'openImage'`
	LogLevel        string `json:'logLevel'`
}

type DBConf struct {
	DbDriver       string `json:'dbDriver'`
	Dsn            string `json:'dsn'`
	DbUserName     string `json:'dbUserName'`
	DbUserPassword string `json:'dbUserPassword'`
	DbHost         string `json:'dbHost'`
	DbPort         string `json:'dbPort'`
	DbName         string `json:'dbName'`
}

type SiteConf struct {
	Nav []Nav `json:'nav'`
}

type Nav struct {
	name string `json:'name'`
	Link string `json:'link'`
}

func GetDBConf() DBConf {
	return dbcfg
}

func GetServeConf() ServeConf {
	return scfg
}

var dbcfg DBConf
var scfg ServeConf

/*
名称：設定ファイル読込処理
概要：各設定ファイルを読み込む
param : 無し
return : 無し
*/
func LoadConfig() {

	raw, err := ioutil.ReadFile(SETTING_DIR_PATH + DB_SETTING_FILE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc Config
	json.Unmarshal(raw, &fc)
	dbcfg = fc.DbInfo

}

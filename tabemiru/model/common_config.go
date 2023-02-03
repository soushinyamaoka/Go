package module

import (
	"encoding/json"
	"fmt"
	"os"
)

/* 名称：サービスベースモジュール
   概要：各サービスモジュールのベースとなるモジュール
*/

// 構造体
type DBconfig struct {
	DbDriver       string `json:'dbDriver'`
	Dsn            string `json:'dsn'`
	DbUserName     string `json:'dbUserName'`
	DbUserPassword string `json:'dbUserPassword'`
	DbHost         string `json:'dbHost'`
	DbPort         string `json:'dbPort'`
	DbName         string `json:'dbName'`
}

func GetDBconfig() DBconfig {
	return cfg
}

var cfg DBconfig

/* 名称：設定ファイル読込処理
   概要：各設定ファイルを読み込む
	 param : 無し
	 return : 無し
*/
func LoadConfig() {
	// DB設定ファイルを読込
	//decode(SETTING_DIR_PATH+DB_SETTING_FILE_NAME, cfg)
	decode("C:/work/PRG/Go/tabemiru/setting/db_config.json", &cfg)
	fmt.Println(cfg)
}

/* 名称：設定ファイル読込処理
   概要：設定ファイルを読み込む
	 param : fn	ファイル名
	 return : 無し
*/
func decode(fn string, v interface{}) {
	// DB設定ファイルを読込
	f, err := os.Open(fn)
	if err != nil {
		// TODO: 例外
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&v)
	fmt.Println(&v)

}

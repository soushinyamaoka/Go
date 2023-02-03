package module

/* 名称：サービスベースモジュール
   概要：各サービスモジュールのベースとなるモジュール
*/
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	con *gorm.DB
}

// type BaseReq struct {
// 	Req http.Request // 種類
// }

type DBService interface {
	Open()
	Close()
}

func (db DB) Open()  {}
func (db DB) Close() {}

// // HomeReqの型定義
// type HomeReq home.HomeReq

// // SearchReqの型定義
// type SearchReq search.SearchReq

/*
DBオープンする
*/
func connectionDB(conf *DBconfig) (*gorm.DB, error) {

	fmt.Println("START:connectionDB")
	// "dsn": "adminhost:rakushirudb@tcp(localhost:3306)/raku"
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		conf.DbUserName,
		conf.DbUserPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	// DB接続
	db, err := gorm.Open(conf.DbDriver, dbConnectInfo)
	fmt.Println("END:connectionDB")

	return db, err
}

/* 名称：サービス取得処理
   概要：リクエストURIに該当するサービスを返す
	 param : http.Request httpリクエスト
	 return : Service サービス
*/
// func GetService(r http.Request) Service {
// 	if URL_PATH_HOOME == r.RequestURI {
// 		return home.HomeReq{Req: r} //
// 	} else if URL_PATH_SEARCH == r.RequestURI {
// 		return search.SearchReq{Req: r} //
// 	} else {
// 		return home.HomeReq{Req: r}
// 	}
// }

// func main() {
// }

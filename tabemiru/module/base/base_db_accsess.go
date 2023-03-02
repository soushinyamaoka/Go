package base

/* 名称：DB接続ベースモジュール
   概要：各DB接続モジュールのベースとなるモジュール
*/
import (
	"fmt"
	c "tabemiru/module/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DBの構造体
type DB struct {
	Con *gorm.DB
}

var err error

type DBService interface {
	Open()
	Close()
}

/*
名称：DB接続処理
概要：DBに接続する
param : 無し
return : error
*/
func (db *DB) Open() error {
	// 設定情報を取得
	conf := c.GetDBConf()
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		conf.DbUserName,
		conf.DbUserPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	// DB接続
	db.Con, err = gorm.Open(conf.DbDriver, dbConnectInfo)
	db.Con.SetLogger(c.GetLogger())
	if err != nil {
		c.GetLog().Error("DB接続に失敗しました。", err.Error())
		return err
	}
	c.GetLog().Info("DB接続に成功しました")
	return err
}

/*
名称：DB切断処理
概要：DBを切断する
param : 無し
return : 無し
*/
func (db DB) Close() error {
	return db.Con.Close()
}

package Module

import (
	"encoding/json"
	"fmt"
	"os"
	Const "rakushiru/src/Module/Const"
	RecipeModel "rakushiru/src/Module/Model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Db struct {
	con *gorm.DB
}

type Result struct {
	Status  int
	Message string
	Err     error
}

type config struct {
	DbDriver       string `json:'dbDriver'`
	Dsn            string `json:'dsn'`
	DbUserName     string `json:'dbUserName'`
	DbUserPassword string `json:'dbUserPassword'`
	DbHost         string `json:'dbHost'`
	DbPort         string `json:'dbPort'`
	DbName         string `json:'dbName'`
}

type Recipes struct {
	RecipeId string `gorm:"recipe_id" json:"RecipeId,omitempty"`
}

type Ingredients struct {
	RecipeId string `gorm:"recipe_id" json:"RecipeId,omitempty"`
}

type Instructions struct {
	RecipeId string `gorm:"recipe_id" json:"RecipeId,omitempty"`
}

/*
設定ファイルを読み込む
*/
func loadConfig() (*config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

/*
DBオープンする
*/
func (db *Db) connectionDB(conf *config) { //(*gorm.DB, error) {

	// "dsn": "adminhost:rakushirudb@tcp(localhost:3306)/raku"
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		conf.DbUserName,
		conf.DbUserPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	var err error
	// DB接続
	db.con, err = gorm.Open(conf.DbDriver, dbConnectInfo)
	if err != nil {
		panic(err.Error())
	}

	// return db, err
}

/*
SELECTを実行する
*/
func (db Db) exeSelectRecipes(model *RecipeModel.Models) {

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	whereRecipes := Recipes{}
	whereIngredients := Ingredients{}
	whereInstructions := Instructions{}
	fmt.Println("SELECT:START")
	if dbResult := db.con.Where(whereRecipes).Find(&model.Recipes); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	if dbResult := db.con.Where(whereIngredients).Find(&model.Ingredients); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	if dbResult := db.con.Where(whereInstructions).Find(&model.Instructions); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("SELECT:END")
	fmt.Println("-----------------------------------")
	// return &recipes
}

/*
SELECTを実行する
*/
func (db Db) exeCountIngredients(where RecipeModel.Ingredients) int {

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	var count int
	fmt.Println("COUNT:START")
	if dbResult := db.con.Model(&Ingredients{}).Where(where).Count(&count); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}

	fmt.Println("COUNT:END")
	fmt.Println("-----------------------------------")
	return count
}

/*
SELECTを実行する
*/
func (db Db) exeCountInstructions(model RecipeModel.Models) int {

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	whereInstructions := Instructions{RecipeId: model.Instructions[0].RecipeId}
	var count int
	fmt.Println("COUNT:START")
	if dbResult := db.con.Model(&Instructions{}).Where(whereInstructions).Count(&count); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}

	fmt.Println("COUNT:END")
	fmt.Println("-----------------------------------")
	return count
}

/*
SELECTを実行する
*/
func (db Db) exeCheckExistRecipes(model *RecipeModel.Models) int {

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	whereRecipes := Recipes{RecipeId: model.Recipes[0].RecipeId}
	var count int
	fmt.Println("COUNT:START")
	if dbResult := db.con.Model(&Recipes{}).Where(whereRecipes).Count(&count); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}

	fmt.Println("COUNT:END")
	fmt.Println("-----------------------------------")
	return count
}

/*
INSERTを実行する
*/
func (db Db) exeInsertRecipeId(model *RecipeModel.Models) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	fmt.Println("INSERT:START")
	if dbResult := db.con.Create(model.Recipes); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return res
}

/*
材料のINSERTを実行する
*/
func (db Db) exeInsertIngredients(model RecipeModel.Ingredients) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	fmt.Println("INSERT:START")
	if dbResult := db.con.Create(model); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return res
}

/*
材料のINSERTを実行する
*/
func (db Db) exeInsertInstructions(model RecipeModel.Instructions) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	fmt.Println("INSERT:START")
	if dbResult := db.con.Create(model); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return res
}

/*
INSERTを実行する
*/
func (db Db) exeInsertRecipe(model *RecipeModel.Models) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	whereRecipes := Recipes{}
	whereRecipes.RecipeId = model.Recipes[0].RecipeId
	fmt.Println("INSERT:START")
	if dbResult := db.con.Where(whereRecipes).Update(&model.Recipes); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	if dbResult := db.con.Create(model.Ingredients); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	if dbResult := db.con.Create(model.Instructions); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return res
}

/*
レシピの更新をする
*/
func (db Db) exeUpdateRecipe(model RecipeModel.Recipes, where Recipes) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	fmt.Println("UPDATE:START")
	if dbResult := db.con.Where(where).Update(&model); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("UPDATE:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
材料の更新をする
*/
func (db Db) exeUpdateIngredients(model RecipeModel.Ingredients, where RecipeModel.Ingredients) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	fmt.Println("UPDATE:START")
	if dbResult := db.con.Where(&where).Update(&model); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("UPDATE:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
手順の更新をする
*/
func (db Db) exeUpdateInstructions(model RecipeModel.Instructions, where Instructions) Result {

	res := Result{}

	defer func() {
		fmt.Println("DBクローズ")
		// DBクローズ
		db.con.Close()
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			panic(errResult)
		}
	}()

	// SQL実行
	fmt.Println("UPDATE:START")
	if dbResult := db.con.Where(&where).Update(&model); dbResult.Error != nil {
		// DBエラーの場合
		panic(Result{Const.STATUS_DB_ERROR, "", dbResult.Error})
	}
	fmt.Println("UPDATE:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
レシピの更新をする
*/
func makeWhereRecipe(model RecipeModel.Models) Result {

	res := Result{}

	whereModel := RecipeModel.Models{}
	whereModel.Recipes[0].RecipeId = model.Recipes[0].RecipeId
	whereIngredients := Ingredients{}
	whereIngredients.RecipeId = model.Recipes[0].RecipeId
	whereIngredients.RecipeId = model.Recipes[0].RecipeId
	whereInstructions := Instructions{}
	whereInstructions.RecipeId = model.Recipes[0].RecipeId
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
レシピの検索をする
*/
func SearchRecipe(model *RecipeModel.Models) Result {

	var res = Result{}

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		panic(Result{Const.STATUS_FILE_LOAD_ERROR, "", err})
	}

	db := Db{}
	// DBオープン
	db.connectionDB(conf)
	// db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		panic(Result{Const.STATUS_DB_ERROR, "", err})
	}

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			res = err.(Result)
			// エラーメッセージのセット
			switch res.Status {
			case Const.STATUS_FILE_LOAD_ERROR:
				res.Message = Const.MSG_FILE_LOAD_ERROR
			case Const.STATUS_DB_ERROR:
				res.Message = Const.MSG_DB_ERROR
			default:

			}
		}
	}()

	// レシピ検索
	db.exeSelectRecipes(model)
	res.Status = Const.STATUS_SUCCESS

	return res
}

/*
レシピの検索をする
*/
func CheckExistRecipe(model *RecipeModel.Models) Result {

	var res = Result{}

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		panic(Result{Const.STATUS_FILE_LOAD_ERROR, "", err})
	}

	db := Db{}
	// DBオープン
	db.connectionDB(conf)
	// db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		panic(Result{Const.STATUS_DB_ERROR, "", err})
	}

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			res = err.(Result)
			// エラーメッセージのセット
			switch res.Status {
			case Const.STATUS_FILE_LOAD_ERROR:
				res.Message = Const.MSG_FILE_LOAD_ERROR
			case Const.STATUS_DB_ERROR:
				res.Message = Const.MSG_DB_ERROR
			default:

			}
		}
	}()

	// レシピ検索
	if db.exeCheckExistRecipes(model) == 0 {
		res.Status = Const.STATUS_DATA_NOT_FIND
	} else {
		res.Status = Const.STATUS_DATA_FIND
	}

	return res
}

/*
材料の登録件数を取得する
*/
func getCountIngredients(model RecipeModel.Models) int {

	var res = Result{}
	var count int = 0

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		panic(Result{Const.STATUS_FILE_LOAD_ERROR, "", err})
	}

	db := Db{}
	// DBオープン
	db.connectionDB(conf)
	// db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		panic(Result{Const.STATUS_DB_ERROR, "", err})
	}

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			res = err.(Result)
			// エラーメッセージのセット
			switch res.Status {
			case Const.STATUS_FILE_LOAD_ERROR:
				res.Message = Const.MSG_FILE_LOAD_ERROR
			case Const.STATUS_DB_ERROR:
				res.Message = Const.MSG_DB_ERROR
			default:

			}
		}
	}()

	// レシピ検索
	count = db.exeCountIngredients(RecipeModel.Ingredients{RecipeId: model.Ingredients[0].RecipeId})

	return count
}

/*
手順の登録件数を取得する
*/
func getCountInstructions(model RecipeModel.Models) int {

	var res = Result{}
	var count int = 0

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		panic(Result{Const.STATUS_FILE_LOAD_ERROR, "", err})
	}

	db := Db{}
	// DBオープン
	db.connectionDB(conf)
	// db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		panic(Result{Const.STATUS_DB_ERROR, "", err})
	}

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			res = err.(Result)
			// エラーメッセージのセット
			switch res.Status {
			case Const.STATUS_FILE_LOAD_ERROR:
				res.Message = Const.MSG_FILE_LOAD_ERROR
			case Const.STATUS_DB_ERROR:
				res.Message = Const.MSG_DB_ERROR
			default:

			}
		}
	}()

	// レシピ検索
	count = db.exeCountInstructions(model)

	return count
}

/*
レシピの登録をする
*/
func SaveRecipe(model RecipeModel.Models) Result {

	var res = Result{}

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		panic(Result{Const.STATUS_FILE_LOAD_ERROR, "", err})
	}

	db := Db{}
	// DBオープン
	db.connectionDB(conf)
	// db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		panic(Result{Const.STATUS_DB_ERROR, "", err})
	}

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			res = err.(Result)
			// エラーメッセージのセット
			switch res.Status {
			case Const.STATUS_FILE_LOAD_ERROR:
				res.Message = Const.MSG_FILE_LOAD_ERROR
			case Const.STATUS_DB_ERROR:
				res.Message = Const.MSG_DB_ERROR
			case Const.STATUS_UNEXPECTED:
				res.Message = Const.MSG_UNEXPECTED_ERROR
			default:

			}
		}
	}()

	count := db.exeCountIngredients(RecipeModel.Ingredients{RecipeId: model.Ingredients[0].RecipeId})

	// レシピデータ更新
	db.exeUpdateRecipe(model.Recipes[0], Recipes{RecipeId: model.Recipes[0].RecipeId})

	// 材料データ更新
	for i := 0; i < count; i++ {
		where := RecipeModel.Ingredients{RecipeId: model.Ingredients[i].RecipeId, OrderNo: i}
		if db.exeCountIngredients(where) > 0 {
			// データが存在する場合
			db.exeUpdateIngredients(model.Ingredients[i], where)

		} else {
			// データが存在しない場合
			db.exeInsertIngredients(model.Ingredients[i])

		}

		// 登録データ > 編集データ
		if db.exeCountIngredients(where) > len(model.Ingredients) {

		}
	}
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
レシピの登録をする
*/
func MakeRecipe(model *RecipeModel.Models) Result {

	var res = Result{}

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			res = err.(Result)
			// エラーメッセージのセット
			switch res.Status {
			case Const.STATUS_FILE_LOAD_ERROR:
				res.Message = Const.MSG_FILE_LOAD_ERROR
			case Const.STATUS_DB_ERROR:
				res.Message = Const.MSG_DB_ERROR
			case Const.STATUS_UNEXPECTED:
				res.Message = Const.MSG_UNEXPECTED_ERROR
			default:

			}
		}
	}()

	// レシピIDを生成
	var i = 0
	for i < 100 {
		random, err := MakeRandomStr(5)
		if err != nil {
			panic(Result{Const.STATUS_UNEXPECTED, "", err})
		}

		if res = CheckExistRecipe(model); res.Status == Const.STATUS_DATA_FIND {
			// データが存在する場合
			i++
		} else {
			model.Recipes[0].RecipeId = random
			break
		}
	}
	if len(model.Recipes[0].RecipeId) > 0 {
		res.Status = Const.STATUS_SUCCESS
	}

	return res

}

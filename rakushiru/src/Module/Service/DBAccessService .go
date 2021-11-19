package Module

import (
	"encoding/json"
	"fmt"
	"os"
	Const "rakushiru/src/Module/Const"
	RecipeModel "rakushiru/src/Module/Model"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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
func connectionDB(conf *config) (*gorm.DB, error) {

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

	return db, err
}

/*
SELECTを実行する
*/
func exeSelRecipes(db *gorm.DB, whereModel RecipeModel.Models) (Result, RecipeModel.Models) {

	model := RecipeModel.Models{}
	whereRecipes := whereModel.Recipes[0]
	whereIngredients := whereModel.Instructions[0]
	whereInstructions := whereModel.Ingredients[0]

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("SELECT:START")
	// レシピ情報を検索
	dbResult := db.Where(whereRecipes).Find(&model.Recipes)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, model
	}

	// 材料情報を検索
	dbResult = db.Where(whereIngredients).Find(&model.Ingredients)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, model
	}

	// 手順情報を検索
	dbResult = db.Where(whereInstructions).Find(&model.Instructions)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, model
	}
	fmt.Println("SELECT:END")
	fmt.Println("-----------------------------------")
	return Result{}, model
}

/*
SELECTを実行する
*/
func exeSelRecipeByStr(db *gorm.DB, where string, key string) (Result, RecipeModel.Models) {

	model := RecipeModel.Models{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("SELECT:START")
	// レシピ情報を検索
	dbResult := db.Where(where, key).Find(&model.Recipes)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, model
	}

	fmt.Println("SELECT:END")
	fmt.Println("-----------------------------------")
	return Result{}, model
}

/*
SELECTを実行する
*/
func exeCountIngredients(db *gorm.DB, where RecipeModel.Ingredients) (Result, int) {

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	var count int
	fmt.Println("IngredientsCOUNT:START")
	dbResult := db.Model(&Ingredients{}).Where(where).Count(&count)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, -1
	}

	fmt.Println("COUNT:END")
	return Result{}, count
}

/*
SELECTを実行する
*/
func exeCountInstructions(db *gorm.DB, model RecipeModel.Instructions) (Result, int) {

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	whereInstructions := Instructions{RecipeId: model.RecipeId}
	var count int
	fmt.Println("InstructionsCOUNT:START")
	dbResult := db.Model(&Instructions{}).Where(whereInstructions).Count(&count)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, -1
	}

	fmt.Println("COUNT:END")
	fmt.Println("-----------------------------------")
	return Result{}, count
}

/*
SELECTを実行する
*/
func exeCheckExistRecipes(db *gorm.DB, where RecipeModel.Recipes) (Result, int) {

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	var count int
	fmt.Println("RecipesCOUNT:START")
	dbResult := db.Model(&Recipes{}).Where(where).Count(&count)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, -1
	}

	fmt.Println("COUNT:END")
	fmt.Println("-----------------------------------")
	return Result{}, count
}

/*
INSERTを実行する
*/
func exeInsRecipe(db *gorm.DB, model RecipeModel.Recipes) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("RecipesINSERT:START")
	fmt.Println(model)
	dbResult := db.Create(model)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return res
}

/*
材料のINSERTを実行する
*/
func exeInsIngredients(db *gorm.DB, model RecipeModel.Ingredients) Result {

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("IngredientsINSERT:START")
	dbResult := db.Create(model)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return Result{}
}

/*
材料のINSERTを実行する
*/
func exeInsInstructions(db *gorm.DB, model RecipeModel.Instructions) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("INSERT:START")
	dbResult := db.Create(model)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("INSERT:END")
	fmt.Println("-----------------------------------")
	return res
}

// /*
// INSERTを実行する
// */
// func exeInsertRecipe(db *gorm.DB, model *RecipeModel.Models) Result {

// 	res := Result{}

// 	defer func() {
// 		// SQL実行時エラーの場合
// 		if errResult := recover(); errResult != nil {
// 			panic(errResult)
// 		}
// 	}()

// 	// SQL実行
// 	whereRecipes := Recipes{}
// 	whereRecipes.RecipeId = model.Recipes[0].RecipeId
// 	fmt.Println("INSERT:START")
// 	// レシピ情報を更新
// 	dbResult := db.Where(whereRecipes).Update(&model.Recipes)
// 	if dbResult.Error != nil {
// 		// DBエラーの場合
// 		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
// 	}
// 	// 材料情報を登録
// 	dbResult = db.Create(model.Ingredients)
// 	if dbResult.Error != nil {
// 		// DBエラーの場合
// 		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
// 	}
// 	// 手順情報を登録
// 	dbResult = db.Create(model.Instructions)
// 	if dbResult.Error != nil {
// 		// DBエラーの場合
// 		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
// 	}
// 	fmt.Println("INSERT:END")
// 	fmt.Println("-----------------------------------")
// 	return res
// }

/*
レシピの更新をする
*/
func exeUpdateRecipe(db *gorm.DB, Recipes RecipeModel.Recipes, where RecipeModel.Recipes) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("RecipeUPDATE:START")
	fmt.Println(Recipes)
	fmt.Println(where)
	dbResult := db.Model(&RecipeModel.Recipes{}).Where(where).Update(&Recipes)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("UPDATE:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
材料の更新をする
*/
func exeUpdIngredients(db *gorm.DB, Ingredients RecipeModel.Ingredients, where RecipeModel.Ingredients) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("Ingredients UPDATE:START")
	fmt.Println(Ingredients)
	fmt.Println(where)
	dbResult := db.Model(&RecipeModel.Ingredients{}).Where(where).Update(&Ingredients)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("UPDATE:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
手順の更新をする
*/
func exeUpdInstructions(db *gorm.DB, Instructions RecipeModel.Instructions, where RecipeModel.Instructions) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("UPDATE:START")
	dbResult := db.Model(&RecipeModel.Instructions{}).Where(where).Update(&Instructions)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("UPDATE:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
材料の削除をする
*/
func exeDelIngredients(db *gorm.DB, Ingredients RecipeModel.Ingredients, where RecipeModel.Ingredients) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("Ingredients Delete:START")
	fmt.Println(Ingredients)
	fmt.Println(where)
	dbResult := db.Model(&RecipeModel.Ingredients{}).Where(where).Delete(&Ingredients)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("Delete:END")
	res.Status = Const.STATUS_SUCCESS

	return res

}

/*
手順の更新をする
*/
func exeDelInstructions(db *gorm.DB, Instructions RecipeModel.Instructions, where RecipeModel.Instructions) Result {

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			db.Close()
		}
	}()

	// SQL実行
	fmt.Println("InstructionsDelete:START")
	dbResult := db.Model(&RecipeModel.Instructions{}).Where(where).Delete(&Instructions)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("Delete:END")
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
func SearchRecipe(keyWord RecipeModel.KeyWord) (Result, RecipeModel.Models) {
	fmt.Println("CALL SearchRecipe")

	var res = Result{}
	model := RecipeModel.Models{}

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return Result{Const.STATUS_FILE_LOAD_ERROR, "", err}, model
	}

	// DBオープン
	db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, model
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

	fmt.Println("検索開始")
	for i := 0; i < len(keyWord.Word); i++ {
		where := Const.COL_TITLE + " LIKE ?"
		key := "%" + keyWord.Word[i] + "%"
		// レシピ検索
		res, model = exeSelRecipeByStr(db, where, key)
		for n := 0; n < len(model.Recipes); n++ {
			fmt.Println(model.Recipes[n].RecipeId)
		}
	}

	return res, model
}

/*
レシピの検索をする
*/
func CheckExistRecipe(model *RecipeModel.Models) Result {

	var res = Result{}
	var count = 0

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return Result{Const.STATUS_FILE_LOAD_ERROR, "", err}
	}

	// DBオープン
	db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}
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
	res, count = exeCheckExistRecipes(db, model.Recipes[0])
	if count == 0 {
		res.Status = Const.STATUS_DATA_NOT_FIND
	} else {
		res.Status = Const.STATUS_DATA_FIND
	}

	return res
}

/*
材料の登録件数を取得する
*/
func getCountIngredients(model RecipeModel.Models) (Result, int) {

	var res = Result{}
	var count int = 0

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return Result{Const.STATUS_FILE_LOAD_ERROR, "", err}, 0
	}

	// DBオープン
	db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, 0
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
	res, count = exeCountIngredients(db, RecipeModel.Ingredients{RecipeId: model.Ingredients[0].RecipeId})

	return res, count
}

/*
手順の登録件数を取得する
*/
func getCountInstructions(model RecipeModel.Models) (Result, int) {

	var res = Result{}
	var count int = 0

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return Result{Const.STATUS_FILE_LOAD_ERROR, "", err}, 0
	}

	db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, 0
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
	res, count = exeCountInstructions(db, model.Instructions[0])

	return res, count
}

/*
レシピの登録をする
*/
func SaveRecipe(model RecipeModel.Models) (Result, string) {

	var res Result = Result{}
	var isNew bool = false
	if len(model.Recipes[0].RecipeId) == 0 {
		isNew = true
	}

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return Result{Const.STATUS_FILE_LOAD_ERROR, "", err}, ""
	}

	// DBオープン
	db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, ""
	}

	fmt.Println("トランザクションを開始します")
	// トランザクション開始
	tx := db.Begin()

	defer func() {

		// エラーが発生した場合
		if err := recover(); err != nil {
			fmt.Println("エラー")
			fmt.Println(err)
			fmt.Println(db)
			tx.Rollback()
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
				res.Status = Const.STATUS_UNEXPECTED
				res.Message = Const.MSG_UNEXPECTED_ERROR
			}
		}
		// DBクローズ
		tx.Close()
	}()

	// レシピIDを生成
	model, res = MakeRecipe(model)

	if res.Err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, ""
	}

	if isNew {
		res = exeInsRecipe(db, model.Recipes[0])

	} else {
		res = exeUpdateRecipe(db, model.Recipes[0], RecipeModel.Recipes{RecipeId: model.Recipes[0].RecipeId})

	}

	if res.Err != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", res.Err}, ""
	}

	// DB登録されている件数を取得
	res, dbCount := exeCountIngredients(db, RecipeModel.Ingredients{RecipeId: model.Recipes[0].RecipeId})
	if res.Err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, ""
	}
	// 画面で設定された件数を取得
	saveCount := len(model.Ingredients)
	// 件数が多い方をDB更新回数として設定
	var updateCount = 0
	if dbCount > saveCount {
		updateCount = dbCount
	} else {
		updateCount = saveCount
	}

	// 材料データ更新
	for i := 0; i < updateCount; i++ {
		// 更新条件を設定
		where := RecipeModel.Ingredients{RecipeId: model.Recipes[0].RecipeId, OrderNo: i}

		// DB更新
		if saveCount >= i {
			// 登録データがある場合
			var ingCount = 0
			res, ingCount = exeCountIngredients(tx, where)
			if ingCount > 0 {
				// データが存在する場合、UPDATEする
				res = exeUpdIngredients(tx, model.Ingredients[i], where)
				if res.Err != nil {
					return Result{Const.STATUS_DB_ERROR, "", err}, ""
				}

			} else if ingCount == 0 {
				// データが存在しない場合、INSERTする
				res = exeInsIngredients(tx, model.Ingredients[i])
				if res.Err != nil {
					return Result{Const.STATUS_DB_ERROR, "", err}, ""
				}

			} else {
				// DBエラー
				return res, ""
			}
		} else {
			//　削除データがある場合
		}
	}

	// DB登録されている件数を取得
	res, dbCount = exeCountInstructions(db, RecipeModel.Instructions{RecipeId: model.Recipes[0].RecipeId})
	if res.Err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, ""
	}
	// 画面で設定された件数を取得
	saveCount = len(model.Instructions)
	// 件数が多い方をDB更新回数として設定
	updateCount = 0
	if dbCount > saveCount {
		updateCount = dbCount
	} else {
		updateCount = saveCount
	}

	// 材料データ更新
	for i := 0; i < updateCount; i++ {
		// 更新条件を設定
		where := RecipeModel.Instructions{RecipeId: model.Recipes[0].RecipeId, OrderNo: i}

		// DB更新
		if saveCount >= i {
			// 登録データがある場合
			var ingCount = 0
			res, ingCount = exeCountInstructions(tx, where)
			if ingCount > 0 {
				// データが存在する場合、UPDATEする
				res = exeUpdInstructions(tx, model.Instructions[i], where)
				if res.Err != nil {
					return Result{Const.STATUS_DB_ERROR, "", err}, ""
				}

			} else if ingCount == 0 {
				// データが存在しない場合、INSERTする
				res = exeInsInstructions(tx, model.Instructions[i])
				if res.Err != nil {
					return Result{Const.STATUS_DB_ERROR, "", err}, ""
				}

			} else {
				// DBエラー
				return res, ""
			}
		} else {
			//　削除データがある場合
		}
	}

	// コミット
	tx.Commit()
	fmt.Println("トランザクションを終了します")

	res.Status = Const.STATUS_SUCCESS

	return res, model.Recipes[0].RecipeId

}

/*
ホーム画面の表示項目を取得する
*/
// func SearchHome(whereModel RecipeModel.Models) (RecipeModel.Models, Result) {
// 	fmt.Println("CALL:SearchHome")
// 	var res Result = Result{}
// 	model := RecipeModel.Models{}
// 	// var model RecipeModel.Models

// 	// DB接続情報取得
// 	conf, err := loadConfig()
// 	if err != nil {
// 		return model, Result{Const.STATUS_FILE_LOAD_ERROR, "", err}
// 	}

// 	// DBオープン
// 	db, err := connectionDB(conf)
// 	// DB接続に失敗した場合
// 	if err != nil {
// 		return model, Result{Const.STATUS_DB_ERROR, "", err}
// 	}

// 	defer func() {

// 		// エラーが発生した場合
// 		if err := recover(); err != nil {
// 			res = err.(Result)
// 			// エラーメッセージのセット
// 			switch res.Status {
// 			case Const.STATUS_FILE_LOAD_ERROR:
// 				res.Message = Const.MSG_FILE_LOAD_ERROR
// 			case Const.STATUS_DB_ERROR:
// 				res.Message = Const.MSG_DB_ERROR
// 			default:

// 			}
// 		}
// 	}()

// 	// レシピ検索
// 	res, model = exeSelectRecipes(db, whereModel)

// 	return model, res
// }

/*
レシピの登録をする
*/
func MakeRecipe(model RecipeModel.Models) (RecipeModel.Models, Result) {

	var res = Result{}
	var recipeId string = ""

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return model, Result{Const.STATUS_FILE_LOAD_ERROR, "", err}
	}

	// DBオープン
	db, err := connectionDB(conf)

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

	fmt.Println("レシピIDを生成")
	// レシピIDを生成
	var i = 0
	for i < 100 {
		random, err := MakeRandomStr(5)
		if err != nil {
			return model, Result{Const.STATUS_UNEXPECTED, "", err}
		}

		// レシピ検索
		res, _ = exeCheckExistRecipes(db, RecipeModel.Recipes{RecipeId: random})
		if res.Status == Const.STATUS_DATA_FIND {
			// データが存在する場合
			i++
		} else {
			recipeId = random
			break
		}
	}
	if len(recipeId) > 0 {
		model = setInitValue(model, recipeId)
	}

	return model, res

}

/*
レシピの登録をする
*/
func setInitValue(model RecipeModel.Models, recipeId string) RecipeModel.Models {

	// 現在時刻取得
	var nowTime time.Time = time.Now()
	if len(recipeId) > 0 {
		// 新規登録の場合、IDと日時を設定

		// レシピ情報
		model.Recipes[0].RecipeId = recipeId
		model.Recipes[0].Image = recipeId + ".jpg"
		model.Recipes[0].CreatedAt = nowTime
		model.Recipes[0].UpdatedAt = nowTime

		// 材料情報
		for i := 0; i < len(model.Ingredients); i++ {
			model.Ingredients[i].RecipeId = recipeId
			model.Ingredients[i].CreatedAt = nowTime
			model.Ingredients[i].UpdatedAt = nowTime
		}

		// 手順情報
		for i := 0; i < len(model.Instructions); i++ {
			model.Instructions[i].RecipeId = recipeId
			model.Instructions[i].CreatedAt = nowTime
			model.Instructions[i].UpdatedAt = nowTime
		}

	} else {
		// 更新の場合

		// レシピ情報
		model.Recipes[0].UpdatedAt = nowTime

		// 材料情報
		for i := 0; i < len(model.Ingredients); i++ {
			model.Ingredients[i].UpdatedAt = nowTime
		}

		// 手順情報
		for i := 0; i < len(model.Instructions); i++ {
			model.Instructions[i].UpdatedAt = nowTime
		}
	}

	return model

}

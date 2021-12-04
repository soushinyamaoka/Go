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
レシピ情報のSELECTを実行する(レシピIDによる検索)
*/
func exeSelRecipes(db *gorm.DB, recipeId string) (Result, RecipeModel.Models) {

	fmt.Println("START:exeSelRecipes")

	model := RecipeModel.Models{}
	// 検索条件のIDを設定
	whereRecipes := Recipes{RecipeId: recipeId}
	whereInstructions := Instructions{RecipeId: recipeId}
	whereIngredients := Ingredients{RecipeId: recipeId}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
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
	fmt.Println("END:exeSelRecipes")

	return Result{}, model
}

/*
新着レシピ情報のSELECTを実行する
*/
func exeSelNewRecipe(db *gorm.DB) (Result, RecipeModel.Models) {
	fmt.Println("START:exeSelNewRecipe")
	model := RecipeModel.Models{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	// レシピ情報を検索
	dbResult := db.Order("updated_at desc").Find(&model.Recipes)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, model
	}

	fmt.Println("END:exeSelNewRecipe")
	return Result{}, model
}

/*
材料情報のSELECTを実行する
*/
func exeCountIngredients(db *gorm.DB, where RecipeModel.Ingredients) (Result, int) {
	fmt.Println("START:exeCountIngredients")

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	var count int
	dbResult := db.Model(&Ingredients{}).Where(where).Count(&count)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, -1
	}

	fmt.Println("END:exeCountIngredients")
	return Result{}, count
}

/*
SELECTを実行する(文字列による検索)
*/
func exeSelRecipeByStr(db *gorm.DB, where string, key string) (Result, RecipeModel.Models) {
	fmt.Println("START:exeSelRecipeByStr")
	model := RecipeModel.Models{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	// レシピ情報を検索
	dbResult := db.Where(where, key).Find(&model.Recipes)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, model
	}

	fmt.Println("END:exeSelRecipeByStr")
	return Result{}, model
}

/*
SELECTを実行する
*/
func exeCountInstructions(db *gorm.DB, where RecipeModel.Instructions) (Result, int) {
	fmt.Println("START:exeCountInstructions")

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	var count int
	dbResult := db.Model(&Instructions{}).Where(where).Count(&count)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, -1
	}

	fmt.Println("END:exeCountInstructions")
	return Result{}, count
}

/*
SELECTを実行する
*/
func exeCheckExistRecipes(db *gorm.DB, where RecipeModel.Recipes) (Result, int) {

	fmt.Println("START:exeCheckExistRecipes")

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	var count int
	dbResult := db.Model(&Recipes{}).Where(where).Count(&count)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}, -1
	}

	fmt.Println("END:exeCheckExistRecipes")
	return Result{}, count
}

/*
INSERTを実行する
*/
func exeInsRecipe(db *gorm.DB, model RecipeModel.Recipes) Result {
	fmt.Println("START:exeInsRecipe")
	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	dbResult := db.Create(model)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("END:exeInsRecipe")
	return res
}

/*
材料のINSERTを実行する
*/
func exeInsIngredients(db *gorm.DB, model RecipeModel.Ingredients) Result {
	fmt.Println("START:exeInsIngredients")
	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	dbResult := db.Create(model)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("END:exeInsIngredients")
	return Result{}
}

/*
材料のINSERTを実行する
*/
func exeInsInstructions(db *gorm.DB, model RecipeModel.Instructions) Result {
	fmt.Println("START:exeInsInstructions")
	fmt.Println(model)

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	dbResult := db.Create(model)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	fmt.Println("END:exeInsInstructions")
	return res
}

/*
レシピの更新をする
*/
func exeUpdateRecipe(db *gorm.DB, Recipes RecipeModel.Recipes, where RecipeModel.Recipes) Result {
	fmt.Println("END:exeUpdateRecipe")

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	dbResult := db.Model(&RecipeModel.Recipes{}).Where(where).Update(&Recipes)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	res.Status = Const.STATUS_SUCCESS

	fmt.Println("END:exeUpdateRecipe")
	return res

}

/*
材料の更新をする
*/
func exeUpdIngredients(db *gorm.DB, Ingredients RecipeModel.Ingredients, where RecipeModel.Ingredients) Result {
	fmt.Println("START:exeUpdIngredients")
	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	dbResult := db.Model(&RecipeModel.Ingredients{}).Where(where).Update(&Ingredients)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	res.Status = Const.STATUS_SUCCESS
	fmt.Println("END:exeUpdIngredients")

	return res

}

/*
手順の更新をする
*/
func exeUpdInstructions(db *gorm.DB, Instructions RecipeModel.Instructions, where RecipeModel.Instructions) Result {
	fmt.Println("START:exeUpdInstructions")
	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	dbResult := db.Model(&RecipeModel.Instructions{}).Where(where).Update(&Instructions)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	res.Status = Const.STATUS_SUCCESS
	fmt.Println("END:exeUpdInstructions")
	return res

}

/*
材料の削除をする
*/
func exeDelIngredients(db *gorm.DB, where RecipeModel.Ingredients) Result {
	fmt.Println("START:exeDelIngredients")

	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	Ingredients := RecipeModel.Ingredients{}

	// SQL実行
	dbResult := db.Model(&RecipeModel.Ingredients{}).Where(where).Delete(&Ingredients)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	res.Status = Const.STATUS_SUCCESS

	fmt.Println("END:exeDelIngredients")
	return res

}

/*
手順の更新をする
*/
func exeDelInstructions(db *gorm.DB, where RecipeModel.Instructions) Result {
	fmt.Println("START:exeDelInstructions")
	res := Result{}

	defer func() {
		// SQL実行時エラーの場合
		if errResult := recover(); errResult != nil {
			fmt.Println("DBクローズ")
			db.Rollback()
			db.Close()
		}
	}()

	// SQL実行
	Instructions := RecipeModel.Instructions{}
	dbResult := db.Model(&RecipeModel.Instructions{}).Where(where).Delete(&Instructions)
	if dbResult.Error != nil {
		// DBエラーの場合
		return Result{Const.STATUS_DB_ERROR, "", dbResult.Error}
	}
	res.Status = Const.STATUS_SUCCESS

	fmt.Println("END:exeDelInstructions")
	return res

}

/*
レシピの更新をする
*/
func makeWhereRecipe(model RecipeModel.Models) Result {
	fmt.Println("END:makeWhereRecipe")

	res := Result{}

	whereModel := RecipeModel.Models{}
	whereModel.Recipes[0].RecipeId = model.Recipes[0].RecipeId
	whereIngredients := Ingredients{}
	whereIngredients.RecipeId = model.Recipes[0].RecipeId
	whereIngredients.RecipeId = model.Recipes[0].RecipeId
	whereInstructions := Instructions{}
	whereInstructions.RecipeId = model.Recipes[0].RecipeId
	res.Status = Const.STATUS_SUCCESS

	fmt.Println("END:makeWhereRecipe")
	return res

}

/*
レシピの検索をする
*/
func OpenRecipeInfo(model RecipeModel.Models) (Result, RecipeModel.Models) {
	fmt.Println("START:OpenRecipeInfo")

	var res = Result{}

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
		db.Close()
	}()

	res, model = exeSelRecipes(db, model.Recipes[0].RecipeId)
	fmt.Println("END:OpenRecipeInfo")

	return res, model
}

/*
レシピの検索をする
*/
func SearchRecipe(keyWord []RecipeModel.Data, isHome bool) (Result, RecipeModel.Models) {
	fmt.Println("START:SearchRecipe")

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
		db.Close()
	}()

	fmt.Println("検索開始")
	var arr []RecipeModel.Recipes
	for i, v := range keyWord {
		fmt.Println(i, v)
		fmt.Println(v)
		where := Const.COL_TITLE + " LIKE ?"
		key := "%" + v.Word + "%"
		// レシピ検索
		var m = RecipeModel.Models{}

		res, m = exeSelRecipeByStr(db, where, key)
		for n, r := range m.Recipes {

			fmt.Println(n, r)
			arr = append(arr, r)
			if isHome && n == 3 {
				break
			}
		}
	}

	model.Recipes = arr
	fmt.Println("END:SearchRecipe")

	return res, model
}

/*
ホーム画面のレシピの検索をする
*/
func SearchHomeRecipe(keyWord []RecipeModel.Data, isHome bool) (Result, RecipeModel.HomeModels) {
	fmt.Println("START:SearchHomeRecipe")

	var res = Result{}
	// model := RecipeModel.Models{}
	hModel := RecipeModel.HomeModels{}

	// DB接続情報取得
	conf, err := loadConfig()
	if err != nil {
		return Result{Const.STATUS_FILE_LOAD_ERROR, "", err}, hModel
	}

	// DBオープン
	db, err := connectionDB(conf)
	// DB接続に失敗した場合
	if err != nil {
		return Result{Const.STATUS_DB_ERROR, "", err}, hModel
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
		db.Close()
	}()

	fmt.Println("検索開始")
	// var arr []RecipeModel.Recipes
	for i, v := range keyWord {
		fmt.Println(i, v)
		fmt.Println(v)
		where := Const.COL_TITLE + " LIKE ?"
		key := "%" + v.Word + "%"
		// レシピ検索
		var m = RecipeModel.Models{}

		res, m = exeSelRecipeByStr(db, where, key)
		// var hm = RecipeModel.HomeModels{}
		// hm.RankModels.Rank1 = append(hm.RankModels.Rank1, m.Recipes[0])
		// var a := append(hm.Recipes.RankModels, m.Recipes)
		for n, r := range m.Recipes {

			fmt.Println(n, r)
			if i == 0 {
				hModel.RankModels.Rank1 = append(hModel.RankModels.Rank1, r)
			} else if i == 1 {
				hModel.RankModels.Rank2 = append(hModel.RankModels.Rank2, r)
			} else if i == 2 {
				hModel.RankModels.Rank3 = append(hModel.RankModels.Rank3, r)
			}
			// arr = append(arr, r)
			if isHome && n == 3 {
				break
			}
		}
	}

	// model.Recipes = arr
	fmt.Println("END:SearchHomeRecipe")

	return res, hModel
}

/*
新着レシピの検索をする
*/
func SearchNewRecipe() (Result, RecipeModel.Models) {
	fmt.Println("START:SearchNewRecipe")

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
		db.Close()
	}()

	fmt.Println("検索開始")
	var arr []RecipeModel.Recipes
	// レシピ検索
	var m = RecipeModel.Models{}
	res, m = exeSelNewRecipe(db)
	for n, r := range m.Recipes {

		fmt.Println(n, r)
		arr = append(arr, r)
	}

	model.Recipes = arr
	fmt.Println("END:SearchNewRecipe")

	return res, model
}

/*
レシピの検索をする
*/
func CheckExistRecipe(model *RecipeModel.Models) Result {
	fmt.Println("START:CheckExistRecipe")

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
		db.Close()
	}()

	// レシピ検索
	res, count = exeCheckExistRecipes(db, model.Recipes[0])
	if count == 0 {
		res.Status = Const.STATUS_DATA_NOT_FIND
	} else {
		res.Status = Const.STATUS_DATA_FIND
	}

	fmt.Println("END:CheckExistRecipe")
	return res
}

/*
材料の登録件数を取得する
*/
func getCountIngredients(model RecipeModel.Models) (Result, int) {
	fmt.Println("END:getCountIngredients")

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
		db.Close()
	}()

	// レシピ検索
	res, count = exeCountIngredients(db, RecipeModel.Ingredients{RecipeId: model.Ingredients[0].RecipeId})

	fmt.Println("START:getCountIngredients")

	return res, count
}

/*
手順の登録件数を取得する
*/
func getCountInstructions(model RecipeModel.Models) (Result, int) {
	fmt.Println("START:getCountInstructions")

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
		db.Close()
	}()

	// レシピ検索
	res, count = exeCountInstructions(db, model.Instructions[0])

	fmt.Println("END:getCountInstructions")
	return res, count
}

/*
レシピの登録をする
*/
func SaveRecipe(model RecipeModel.Models) (Result, string) {
	fmt.Println("START:SaveRecipe")

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
		fmt.Println("でぃふぁー")

		// エラーが発生した場合
		if err := recover(); err != nil {
			fmt.Println("エラーリカバー：ロールバック")
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
		} else if res.Status != Const.STATUS_SUCCESS {
			fmt.Println("エラー：ロールバック")
			tx.Rollback()
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
		fmt.Println("トランザクションクローズ")
		tx.Close()
	}()

	// レシピIDを生成
	if isNew {
		model, res = MakeRecipe(model)
	} else {
		model = setInitValue(model, "")
	}

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
		where := RecipeModel.Ingredients{RecipeId: model.Recipes[0].RecipeId, OrderNo: i + 1}

		// DB更新
		if saveCount >= i {
			// 登録データがある場合
			var ingCount = 0
			res, ingCount = exeCountIngredients(tx, where)
			if ingCount > 0 {
				// データが存在する場合、UPDATEする
				if saveCount > i {
					// 更新の場合
					res = exeUpdIngredients(tx, model.Ingredients[i], where)
				} else {
					// 削除の場合
					// res = exeDelIngredients(tx, model.Ingredients[i], where)
					res = exeDelIngredients(tx, where)
				}
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
		where := RecipeModel.Instructions{RecipeId: model.Recipes[0].RecipeId, OrderNo: i + 1}

		// DB更新
		if saveCount >= i {
			// 登録データがある場合
			var ingCount = 0
			res, ingCount = exeCountInstructions(tx, where)
			if ingCount > 0 {
				// データが存在する場合、UPDATEする
				if saveCount > i {
					// 更新の場合
					res = exeUpdInstructions(tx, model.Instructions[i], where)
				} else {
					// 削除の場合
					res = exeDelInstructions(tx, where)
				}
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
	fmt.Println("END:SaveRecipe")

	return res, model.Recipes[0].RecipeId

}

/*
レシピの登録をする
*/
func MakeRecipe(model RecipeModel.Models) (RecipeModel.Models, Result) {
	fmt.Println("START:MakeRecipe")

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
		db.Close()
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
	fmt.Println("END:MakeRecipe")

	return model, res

}

/*
レシピの登録をする
*/
func setInitValue(model RecipeModel.Models, recipeId string) RecipeModel.Models {
	fmt.Println("START:setInitValue")

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
		if len(model.Ingredients) > 0 {
			for i := 0; i < len(model.Ingredients); i++ {
				model.Ingredients[i].RecipeId = recipeId
				model.Ingredients[i].CreatedAt = nowTime
				model.Ingredients[i].UpdatedAt = nowTime
			}
		} else {
			fmt.Println("材料情報はありません")
		}

		// 手順情報
		if len(model.Instructions) > 0 {
			for i := 0; i < len(model.Instructions); i++ {
				model.Instructions[i].RecipeId = recipeId
				model.Instructions[i].CreatedAt = nowTime
				model.Instructions[i].UpdatedAt = nowTime
			}
		} else {
			fmt.Println("手順情報はありません")
		}

	} else {
		// 更新の場合

		// レシピ情報
		model.Recipes[0].UpdatedAt = nowTime

		// 材料情報
		for i := 0; i < len(model.Ingredients); i++ {
			if model.Ingredients[i].CreatedAt.IsZero() {
				fmt.Println("0です")
				model.Ingredients[i].CreatedAt = nowTime
			}
			model.Ingredients[i].UpdatedAt = nowTime
		}

		// 手順情報
		for i := 0; i < len(model.Instructions); i++ {
			if model.Instructions[i].CreatedAt.IsZero() {
				fmt.Println("0です")
				model.Instructions[i].CreatedAt = nowTime
			}
			model.Instructions[i].UpdatedAt = nowTime
		}
	}

	fmt.Println("END:setInitValue")
	return model

}

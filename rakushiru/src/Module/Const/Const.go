package Const

const (
	// ステータスコード
	STATUS_SUCCESS         = 0
	STATUS_DATA_FIND       = 1
	STATUS_DATA_NOT_FIND   = 2
	STATUS_DB_ERROR        = 100
	STATUS_FILE_LOAD_ERROR = 300
	STATUS_UNEXPECTED      = 999
	// 結果メッセージ
	MSG_FILE_LOAD_ERROR  = "ファイル読込失敗"
	MSG_DB_ERROR         = "DBエラー"
	MSG_UNEXPECTED_ERROR = "何らかのエラー"
	// リクエストコード
	REQ_SAVE_RECIPE   = "saveRecipe"
	REQ_SEARCH_RECIPE = "searchRecipe"
	// カラム名
	COL_TITLE = "title"
)

func getMaterialStub() {
	// RecipeModel.Foo()
	//ping = RecipeModel.Ping{1, "ok"}
}

package common

/* 名称：サービスベースモジュール
   概要：各サービスモジュールのベースとなるモジュール
*/
const (
	SETTING_DIR_PATH     = "setting/"
	DB_SETTING_FILE_NAME = "config.json"

	URL_PATH_HOOME  = "/Home"
	URL_PATH_SEARCH = "/Search"
)

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
	REQ_SAVE_RECIPE       = "saveRecipe"
	REQ_SEARCH_RECIPE     = "searchRecipe"
	REQ_SEARCH_NEW_RECIPE = "searchNewRecipe"
	// カラム名
	COL_TITLE = "title"
)

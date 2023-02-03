package module

/* 名称：サービスベースモジュール
   概要：各サービスモジュールのベースとなるモジュール
*/
import (
	"net/http"
	home "tabemiru/model/home"
	search "tabemiru/model/search"
)

// 構造体
type BaseReq struct {
	Req http.Request // 種類
}

type Service interface {
	Sample()
	Excute()
}

func (r BaseReq) Excute() {}
func (r BaseReq) Sample() {}

// HomeReqの型定義
type HomeReq home.HomeReq

// SearchReqの型定義
type SearchReq search.SearchReq

/* 名称：サービス取得処理
   概要：リクエストURIに該当するサービスを返す
	 param : http.Request httpリクエスト
	 return : Service サービス
*/
func GetService(r http.Request) Service {
	if URL_PATH_HOOME == r.RequestURI {
		return home.HomeReq{Req: r} //
	} else if URL_PATH_SEARCH == r.RequestURI {
		return search.SearchReq{Req: r} //
	} else {
		return home.HomeReq{Req: r}
	}
}

func main() {
}

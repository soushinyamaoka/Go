package search

import (
	"fmt"
	"net/http"
)

// 構造体
type SearchReq struct {
	Req http.Request // 種類
}

type Service interface {
	Sample()
	Excute()
}

func (r SearchReq) Excute() {
	fmt.Println("Search.Excute")
}
func (r SearchReq) Sample() {
	fmt.Println("Search.Sample")
}

func main() {
}

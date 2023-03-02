package common

/* 名称：サービエラーモジュール
   概要：発生したエラーを処理する
*/
type MyError struct {
	Code       string
	Msg        string
	StackTrace string
}

// // Error error interfaceを実装
// func (me *MyError) Error() string {
//   return fmt.Sprintf("my error: code[%s], message[%s]", me.Code, me.Msg)
// }

// // New コンストラクタ
// func New(code string, msg string) *MyError {
//   stack := zap.Stack("").String
//   return &MyError{
//     Code:       code,
//     Msg:        msg,
//     StackTrace: stack,
//   }
// }

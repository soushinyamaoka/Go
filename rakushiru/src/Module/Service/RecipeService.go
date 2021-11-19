package Module

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func FileUpload(file multipart.File, recipeId string) {
	// ファイル名生成
	uploadedFileName := recipeId + ".jpg"
	//アップロードされたファイルを置くパスを設定
	imagePath := "image/" + uploadedFileName

	//imagePathにアップロードされたファイルを保存
	saveImage, err := os.Create(imagePath)

	if err != nil {
		fmt.Println("ファイルの確保失敗")
	}

	//保存用ファイルにアップロードされたファイルを書き込む
	_, err = io.Copy(saveImage, file)
	if err != nil {
		fmt.Println("アップロードしたファイルの書き込み失敗")
	}

	//saveImageとfileを最後に閉じる
	defer saveImage.Close()
	defer file.Close()
}

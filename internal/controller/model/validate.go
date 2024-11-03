package model

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

// ValidateTranscribe
func ValidateTranscribe(c echo.Context) (string, error) {
	// input
	// 音声ファイルを取得
	file, err := c.FormFile("audio")
	if err != nil {
		return "", err
	}
	// アップロードされたファイルを開く
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 一時ファイルを作成
	tempFile, err := os.CreateTemp("", "audio-*.wav") // 拡張子は必要に応じて変更
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name()) // 処理が終わったら一時ファイルを削除

	// 一時ファイルにデータを書き込む
	if _, err := io.Copy(tempFile, src); err != nil {
		return "", err
	}

	// 一時ファイルのパスを取得
	tempFilePath := tempFile.Name()

	return tempFilePath, nil
}

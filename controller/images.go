package controller

import (
	"encoding/base64"
	"fmt"
	"onlinebc_admin/model/ftp"
	"path/filepath"
	"strings"
)

// saveImage сохраняет изображение и возвращает URI изображения и его иконки
func saveImage(parentID int, fileName string, base64string string) (imageURI string, thumbURI string) {
	imageBytes, err := base64.StdEncoding.DecodeString(base64string)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	resizedImage := resizeImage(imageBytes)
	thumb := makeThumb(imageBytes)
	dir := getDirName(parentID)
	ext := filepath.Ext(fileName)
	thumbName := strings.TrimSuffix(fileName, ext) + "_thumb" + ext

	imageURI = ftp.Save(dir, fileName, resizedImage)
	thumbURI = ftp.Save(dir, thumbName, thumb)

	return
}

func getDirName(parentID int) (dirname string) {
	dirname = fmt.Sprintf("uploaded_images/%08d", parentID)
	return
}

func resizeImage(imageBytes []byte) []byte {
	return imageBytes
}

func makeThumb(imageBytes []byte) []byte {
	return imageBytes
}

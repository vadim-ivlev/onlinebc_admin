// Package img - функции для декодирования, масштабирования и сохранения изображений в файл.
package img

import (
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"strings"

	"github.com/disintegration/imaging"
)

const (
	// директория для временного хранения загруженных фото
	uploadsTempDir = "uploads_temp/"
	// максимальная ширина сохраняемого изображения
	maxImageWidth = 1440
	// размеры иконки
	thumbWidth  = 310
	thumbHeight = 206
)

var (
	validExtensions = map[string]bool{".jpeg": true, ".jpg": true, ".png": true, ".gif": true}
)

// SaveImage сохраняет изображение и возвращает URI изображения и URI иконки
func SaveImage(parentID int, fileName string, base64string string) (imageURI string, thumbURI string) {
	ext := filepath.Ext(fileName)
	if !validExtensions[ext] {
		return
	}
	thumbName := strings.TrimSuffix(fileName, ext) + "_thumb" + ext
	im, _, _ := imageFromString(base64string)
	if im != nil {
		err := os.MkdirAll(uploadsTempDir, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
		imageURI = resizeAndSave(im, uploadsTempDir+fileName)
		thumbURI = thumbAndSave(im, uploadsTempDir+thumbName)
	}
	return
}

// resizeAndSave масштабирует изображение если необходимо и сохраняет его в файл
func resizeAndSave(im image.Image, filePath string) string {
	dst := im
	b := dst.Bounds()
	width := b.Dx()
	height := b.Dy()
	if width > maxImageWidth {
		dst = imaging.Resize(im, maxImageWidth, height*maxImageWidth/width, imaging.Lanczos)
	}
	return saveImageToFile(dst, filePath)
}

// thumbAndSave генерирует уменьшенное изображение и сохраняет его в файл
func thumbAndSave(im image.Image, filePath string) string {
	b := im.Bounds()
	width := b.Dy()
	height := b.Dy()
	anchor := imaging.Top
	if width/height > thumbWidth/thumbHeight {
		anchor = imaging.Center
	}
	dst := imaging.Fill(im, thumbWidth, thumbHeight, anchor, imaging.Lanczos)
	return saveImageToFile(dst, filePath)
}

// Сохранение файла
func saveImageToFile(dst image.Image, filePath string) string {
	err := imaging.Save(dst, filePath)
	if err != nil {
		fmt.Println("ERROR saveImageToFile ****************************************")
		fmt.Println(err.Error())
		return ""
	}
	return filePath
}

// imageFromString декодирует строку base64 в структуру image.Image
func imageFromString(imageString string) (image.Image, string, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageString))
	im, format, err := image.Decode(reader)
	if err != nil {
		fmt.Println("ERR:", err.Error())
		return nil, "", err
	}
	return im, format, nil
}

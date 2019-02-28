package img

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"os"
	"sportbc_admin/config"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

const (
	// длинна имени для фото
	LenPhotoName = 15
	// корневая директория для загрузки фото
	frontPathUpload = "/uploads"
)

var (
	// разрешенные для загрузки расширения
	allowePhotodExt = [5]string{".jpg", ".jpeg", ".png", ".gif"}
	// ErrEmptyDirPath в аргументах передана пустая директория
	ErrEmptyDirPath = errors.New("Path to directory is empty")
)

// Resize по заданным параметрам
func Resize(filePath string, w, h int) error {
	src, err := imaging.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "Failed to open image")
	}
	dst := imaging.Resize(src, w, h, imaging.Lanczos)
	// Сохранение файла
	err = imaging.Save(dst, filePath)
	if err != nil {
		return errors.Wrap(err, "Failed to save image")
	}
	return nil
}

// Fill по заданным параметрам
func Fill(filePath string, fileDest string, w, h int) error {
	src, err := imaging.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "Failed to open image")
	}
	dst := imaging.Fill(src, w, h, imaging.Center, imaging.Lanczos)
	// Сохранение файла
	err = imaging.Save(dst, fileDest)
	if err != nil {
		return errors.Wrap(err, "Failed to save image")
	}
	return nil
}

// Dimension получение размеров у фото
func Dimension(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return 0, 0, errors.Wrap(err, "Dimension Open file")
	}

	i, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Dimension DecodeConfig")
	}
	return i.Width, i.Height, nil
}

// DimensionFile получение размеров файла (без загрузки на сервер)
func DimensionFile(file []byte) (int, int, error) {
	r := bytes.NewReader(file)
	i, _, err := image.DecodeConfig(r)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Dimension DecodeConfig")
	}
	// fmt.Println(i.Width, i.Height)
	return i.Width, i.Height, nil
}

// DeleteFile удаление фото с сервера
func DeleteFile(filePath string) error {
	// удаляем если файл существует
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveDir удаление директории с фото и превью
func RemoveDir(dir string) error {
	// проверка на пустую директорию, иначе есть шанс удалить все от корня проекта
	if strings.TrimSpace(dir) == "" {
		return errors.Wrap(ErrEmptyDirPath, "Failed to remove dir "+dir)
	}
	// удаляем если файл существует
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckExt проверка на допустимые расширения для фото
func CheckExt(ext string) bool {
	for _, v := range allowePhotodExt {
		if v == ext {
			return true
		}
	}
	return false
}

// GenPhotoPath генерация пути до фотографии
func GenPhotoPath(name string) string {
	path := frontPathUpload + "/" + time.Now().Format("2006/01/02") + "/" + name
	// Если директории не существует - создаем ее во временной
	MkdirPath(config.APP.TempPATH + path)
	return path
}

// MkdirPath создание директории хранения фото, если ее не существует
func MkdirPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}

// GenFileName генерация уникального имени для фотографии
func GenFileName(fileName string) string {
	s := strconv.FormatInt(time.Now().UnixNano(), 10) + fileName
	h := sha256.New()
	h.Write([]byte(s))
	name := hex.EncodeToString(h.Sum(nil))
	name = name[0:LenPhotoName]
	return name
}

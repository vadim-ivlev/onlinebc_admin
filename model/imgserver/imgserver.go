// Package imgserver содержит функции для перемещения файлов изображений
// из временной директории на удаленный сервер по ssh.
// Используются команды ssh и scp, предположительно доступные
// в операционной системе.
package imgserver

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	// директория для временного хранения загруженных фото
	tempUploadPath = "/uploads"
)

// GenPhotoDir генерация директории для сохранения фотографии
func GenPhotoDir() string {
	return tempUploadPath + "/" + time.Now().Format("2006/01/02")
}

// MoveFileToImageServer перемещает файл  на удаленный сервер
func MoveFileToImageServer(filePath string) string {
	fileName := filepath.Base(filePath)
	photoDir := GenPhotoDir()
	destDir := params.Uploaddir + photoDir

	cmdMkdir := fmt.Sprintf("sshpass -p %s ssh -q -o StrictHostKeyChecking=no -p %s %s@%s mkdir -p %s; exit", params.Password, params.Port, params.User, params.Host, destDir)
	cmdCopyFile := fmt.Sprintf("sshpass -p %s scp -o StrictHostKeyChecking=no -P %s %s %s@%s:%s/", params.Password, params.Port, filePath, params.User, params.Host, destDir)
	cmdRemoveFile := fmt.Sprintf("rm %s", filePath)

	bash(cmdMkdir)
	bash(cmdCopyFile)
	bash(cmdRemoveFile)

	return photoDir + "/" + fileName
}

func bash(cmdString string) (errMessage string) {
	cmd := exec.Command("bash")
	cmdWriter, _ := cmd.StdinPipe()
	errReader, _ := cmd.StderrPipe()
	cmd.Start()
	cmdWriter.Write([]byte(cmdString + "\n"))
	cmdWriter.Write([]byte("exit" + "\n"))
	errBytes, _ := ioutil.ReadAll(errReader)
	cmd.Wait()
	errMessage = string(errBytes)
	if errMessage != "" {
		fmt.Println(cmdString, "ERR:", errMessage)
	}
	return
}

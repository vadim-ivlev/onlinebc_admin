// Package imgserver содержит функции для перемещения файлов изображений
// из временной директории на удаленный сервер по ssh.
// Используются команды ssh и scp, которые должны быть доступны
// в операционной системе.
package imgserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Params - общие параметры хранимые в YAML
var Params struct {
	Host      string `yaml:"host"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Port      string `yaml:"port"`
	Remotedir string `yaml:"remotedir"`
	Localdir  string `yaml:"localdir"`
}

// ReadConfig читает YAML
func ReadConfig(fileName string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := yaml.Unmarshal(yamlFile, &Params); err != nil {
		log.Println(err.Error())
	}
}

// Bash - исполняет команду bash, возвращает строку stderr.
func Bash(cmdString string) (errMessage string) {
	cmd := exec.Command("bash")
	cmdWriter, _ := cmd.StdinPipe()
	errReader, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Println("BASH cmd.Start(): ", err)
	}

	_, err = cmdWriter.Write([]byte(cmdString + "\n"))
	if err != nil {
		log.Println("BASH Write(): ", err)
	}

	_, err = cmdWriter.Write([]byte("exit" + "\n"))
	if err != nil {
		log.Println("BASH Write(): ", err)
	}

	errBytes, _ := ioutil.ReadAll(errReader)
	err = cmd.Wait()
	if err != nil {
		log.Println("BASH Write(): ", err)
	}

	errMessage = string(errBytes)
	if errMessage != "" {
		fmt.Println(cmdString, "\nBASH MESSAGE:\n", errMessage)
	}
	return
}

// CopyTempFilesToServer - копирует содержимое временной директории загрузки на удаленный сервер.
func CopyTempFilesToServer() string {
	cmd := fmt.Sprintf("sshpass -p %s scp -r -q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -P %s %s/** %s@%s:%s/", Params.Password, Params.Port, Params.Localdir, Params.User, Params.Host, Params.Remotedir)
	// fmt.Println(cmd)
	return Bash(cmd)
}

// RemoveEmptyDirectories = рекурсивно удаляет пустые субдиректории из
// временной директории загрузки
func RemoveEmptyDirectories() string {
	cmd := fmt.Sprintf("find %s/* -type d -empty -delete", Params.Localdir)

	// работает в mac OS, и Linux
	//cmd := fmt.Sprintf("find %s/* -type d -exec rmdir -p {} + 2>/dev/null", Params.Localdir)
	return Bash(cmd)
}

// TrimLocaldir - удаляет префикс временной директории загрузки из пути файла
func TrimLocaldir(path string) string {
	return strings.TrimPrefix(path, Params.Localdir)
}

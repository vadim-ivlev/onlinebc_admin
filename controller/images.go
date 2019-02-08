package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func saveImage(parentID int, fileName string, base64string string) (path string) {
	decoded, err := base64.StdEncoding.DecodeString(base64string)
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}
	dirname := fmt.Sprintf("uploaded_images/%08d", parentID)
	// newpath := filepath.Join(".", "public")
	err2 := os.MkdirAll(dirname, os.ModePerm)
	if err2 != nil {
		panic(err2)
	}

	path = fmt.Sprintf("%s/%s", dirname, fileName)
	fmt.Println(path)
	err1 := ioutil.WriteFile(path, decoded, 0644)
	if err1 != nil {
		panic(err1)
	}
	return
}

package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func saveImage(base64string string) (path string) {
	decoded, err := base64.StdEncoding.DecodeString(base64string)
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}
	path = "blabla.jpg"
	err1 := ioutil.WriteFile(path, decoded, 0644)
	if err1 != nil {
		panic(err1)
	}

	return
}

package embeds

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var textBytes []byte
var jsonBytes []byte

func ReadTextFile(fileName string) []byte {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

// зачитываем текстовые файлы
func TestMain(m *testing.M) {
	textBytes = ReadTextFile("post-text.txt")
	jsonBytes = ReadTextFile("_broadcast354.json")
	os.Exit(m.Run())
}

// Чистим текст и строим список эмбедов
func Test_GetClearTextAndWidgets(m *testing.T) {
	txt, ws := GetClearTextAndWidgets(string(textBytes))
	fmt.Println(txt, strings.Repeat("=", 100)+"\n", ws)
	for i, w := range ws {
		fmt.Println(i)
		for k, v := range w {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
}

// преобразуем финальный JSON
func Test_AmendPostsAndAnswers(m *testing.T) {
	fmt.Println(string(AmendPostsAndAnswers(jsonBytes)))
}

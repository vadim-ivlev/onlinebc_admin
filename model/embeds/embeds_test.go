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

func TestMain(m *testing.M) {
	fmt.Println("Тесты Embeds ******************************************************")

	textBytes = ReadTextFile("post-text.txt")
	jsonBytes = ReadTextFile("_broadcast354.json")

	os.Exit(m.Run())
}

func Test_GetSocialsJSON(m *testing.T) {

	// jsonBytes := ReadTextFile("_broadcast354.json")
	// fmt.Println(string(jsonBytes))

	txt, json := GetClearTextAndWidgets(string(textBytes))
	fmt.Println(txt, strings.Repeat("=", 100)+"\n", json)
}

func Test_traversePosts(m *testing.T) {
	amendPostsAndAnswers(jsonBytes)
}

// embeds Парсинг виджетов в посте трансляции.
// Из исходного текста удаляем эмбеды заменяя их на пробелы.
// https://git.rgwork.ru/masterback/onlinebc_admin/issues/27

package embeds

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	// "github.com/tidwall/gjson"
)

// регулярные выражения фрагментов для удаления из текста
var fragmentsToClear = map[string]*regexp.Regexp{
	"vkpost":     regexp.MustCompile(`(?s)<div.*?vk_post_.*?</div>`),
	"fbpost":     regexp.MustCompile(`(?s)<div.*?class="fb-post".*?</div>`),
	"script":     regexp.MustCompile(`(?s)<script.*?</script>`),
	"blockquote": regexp.MustCompile(`(?s)<blockquote.*?</blockquote>`),
	"iframe":     regexp.MustCompile(`(?s)<iframe.*?</iframe>`),
}

// регулярное выражение для обнаружения виджета в тексте и функция для генерации  заготовки JSON.
type soc struct {
	re  *regexp.Regexp
	gen func(string) map[string]string
}

// перечень типов виджетов для обработки
var widgets = map[string]soc{
	"Youtube": soc{
		re: regexp.MustCompile(`https://www.youtube.com/embed/[^"]*`),
		gen: func(chunk string) map[string]string {
			return map[string]string{
				"type":         "youtube",
				"data-videoid": getLastPart(chunk),
			}
		},
	},
	"Instagram": soc{
		re: regexp.MustCompile(`https://www.instagram.com/p/[^"]*`),
		gen: func(chunk string) map[string]string {
			return map[string]string{
				"type":           "instagram",
				"data-shortcode": getLastPart(chunk),
			}
		},
	},
	"Twitter": soc{
		re: regexp.MustCompile(`https://twitter.com/[^/]+/status/[^?]*`),
		gen: func(chunk string) map[string]string {
			return map[string]string{
				"type":         "twitter",
				"data-tweetid": getLastPart(chunk),
			}
		},
	},
	"facebook": soc{
		re: regexp.MustCompile(`https://www.facebook.com/[^"]*`),
		gen: func(chunk string) map[string]string {
			return map[string]string{
				"type":      "facebook",
				"data-href": chunk,
			}
		},
	},
	"VK": soc{
		re: regexp.MustCompile(`VK.Widgets.Post\([^)]*`),
		gen: func(chunk string) map[string]string {
			return map[string]string{
				"type":           "vk",
				"data-embedtype": "embedtype",
				"data-owner-id":  "oid",
				"data-post-id":   "pid",
				"data-hash":      "hash",
			}
		},
	},
}

func getLastPart(ss string) string {
	s := strings.TrimSuffix(ss, "/")
	a := strings.Split(s, "/")
	return a[len(a)-1]
}

// GetClearTextAndWidgets Возвращает очищенный текст и социальные вставки найденные в тексте.
// https://git.rgwork.ru/masterback/onlinebc_admin/issues/27
func GetClearTextAndWidgets(text string) (strippedText string, widgetsJSON []map[string]string) {
	// накопительный массив для записей о виджета. FIXME:?
	arr := make([]map[string]string, 0)

	// для всех видов виджетов
	for _, w := range widgets {
		// выдираем из текста характерные кусочки
		extracts := w.re.FindAllString(text, -1)
		// для каждого кусочка генерируем структуру и добавляем ее в массив
		for _, extract := range extracts {
			m := w.gen(extract)
			arr = append(arr, m)
		}
	}

	// сериализуем массив
	// bytes, err := json.MarshalIndent(arr, "", "  ")
	// if err != nil {
	// 	log.Println(err)
	// }
	// widgetsJSON = string(bytes)
	widgetsJSON = arr
	strippedText = ClearText(text)
	return
}

// ClearText - возвращает текст очищенный от социальных вставок.
func ClearText(text string) string {
	s := text
	for k, re := range fragmentsToClear {
		s = re.ReplaceAllString(s, strings.ToUpper(k))
	}
	return s
}

// const json1 = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

// func aaa() {
// 	value := gjson.Get(json1, "name.last")

// 	println(value.String())
// }

func amendPostsAndAnswers(jsonBytes []byte) string {
	var broadcasts []Broadcast
	err := json.Unmarshal(jsonBytes, &broadcasts)
	if err != nil {
		fmt.Println("ERR", err)
	}

	for i, broadcast := range broadcasts {
		for j, post := range broadcast.Posts {
			txt, ws := GetClearTextAndWidgets(post.PostsText)
			// post.PostsClearText = txt
			// post.PostsSocials = jsonText

			broadcasts[i].Posts[j].PostsClearText = txt
			broadcasts[i].Posts[j].PostsSocials = ws
			//TODO: answers
		}
	}

	// сериализуем массив
	bytes, err := json.MarshalIndent(broadcasts, "", "  ")
	if err != nil {
		log.Println("ERR serial:", err)
	}
	finalText := string(bytes)
	fmt.Println(finalText)
	return finalText

}

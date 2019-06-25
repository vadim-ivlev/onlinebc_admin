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
)

// регулярные выражения фрагментов для удаления из текста

var fragmentsToClear = []*regexp.Regexp{
	regexp.MustCompile(`(?s)<script.*?</script>`),
	regexp.MustCompile(`(?s)<iframe.*?</iframe>`),
	regexp.MustCompile(`(?s)<blockquote.*?</blockquote>`),
	regexp.MustCompile(`(?s)<div.*?vk_post_.*?</div>`),
	regexp.MustCompile(`(?s)<div.*?class="fb-post".*?</div>`),
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
		re: regexp.MustCompile(`data-instgrm-permalink="https://www.instagram.com/p/[^"]*`),
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
			s := strings.TrimPrefix(chunk, "VK.Widgets.Post(")
			args := strings.Split(s, ",")
			firstArgParts := strings.Split(args[0], "_")
			embedType := firstArgParts[1]
			ownerID := strings.Trim(args[1], " ")
			postID := strings.Trim(args[2], " ")
			hashQuoted := strings.Trim(args[3], " ")
			hash := strings.Trim(hashQuoted, "'")
			return map[string]string{
				"type":           "vk",
				"data-embedtype": embedType,
				"data-owner-id":  ownerID,
				"data-post-id":   postID,
				"data-hash":      hash,
			}
		},
	},
}

// getLastPart возвращает последнюю часть пути
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
	widgetsJSON = arr
	strippedText = ClearText(text)
	return
}

// ClearText - возвращает текст очищенный от социальных вставок.
func ClearText(text string) string {
	s := text
	for _, re := range fragmentsToClear {
		s = re.ReplaceAllString(s, "")
	}
	return s
}

// AmendPostsAndAnswers для каждого поста и ответа к посту добавляем вычисляемые поля.
// Основная функция пакета.
func AmendPostsAndAnswers(jsonBytes []byte) []byte {

	// десериализуем текст в структуру
	var broadcasts []Broadcast
	err := json.Unmarshal(jsonBytes, &broadcasts)
	if err != nil {
		fmt.Println("ERR", err)
	}

	// для каждого поста и ответа к посту добавляем вычисляемые поля
	for i, broadcast := range broadcasts {

		// для каждого поста добавляем вычисляемые поля
		for j, post := range broadcast.Posts {
			txt, ws := GetClearTextAndWidgets(post.PostsText)
			broadcasts[i].Posts[j].PostsClearText = txt
			broadcasts[i].Posts[j].PostsEmbeds = ws

			// для каждого ответа к посту добавляем вычисляемые поля
			for k, answer := range post.PostsAnswers {
				txt, ws := GetClearTextAndWidgets(answer.PostsAnswerText)
				broadcasts[i].Posts[j].PostsAnswers[k].PostsAnswerClearText = txt
				broadcasts[i].Posts[j].PostsAnswers[k].PostsAnswerEmbeds = ws
			}
		}
	}

	// сериализуем структуру обратно в текст
	bytes, err := json.MarshalIndent(broadcasts, "", "  ")
	if err != nil {
		log.Println("ERR serial:", err)
	}
	return bytes
}

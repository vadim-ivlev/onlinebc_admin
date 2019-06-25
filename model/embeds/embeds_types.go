package embeds

// Типы используемые для десериализации/сериализации JSON перед отдачей его клиенту

type Broadcasts []Broadcast

type Broadcast struct {
	ID           int     `json:"id"`
	Posts        []Posts `json:"posts"`
	Title        string  `json:"title"`
	IsDiary      int     `json:"is_diary"`
	IsEnded      int     `json:"is_ended"`
	LinkImg      string  `json:"link_img"`
	ShowDate     int     `json:"show_date"`
	ShowTime     int     `json:"show_time"`
	TimeBegin    int     `json:"time_begin"`
	DiaryAuthor  string  `json:"diary_author"`
	LinkArticle  string  `json:"link_article"`
	TimeCreated  int     `json:"time_created"`
	GroupsCreate int     `json:"groups_create"`
	ShowMainPage int     `json:"show_main_page"`
}
type Thumbs struct {
	Type     string `json:"type"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filepath string `json:"filepath"`
}
type PostsImages struct {
	ID       int         `json:"id"`
	Width    interface{} `json:"width"`
	Height   interface{} `json:"height"`
	Source   string      `json:"source"`
	Thumbs   []Thumbs    `json:"thumbs"`
	PostID   int         `json:"post_id"`
	Filepath string      `json:"filepath"`
}
type PostsAnswers struct {
	ID                   int         `json:"id"`
	IDParent             int         `json:"id_parent"`
	HasBigImg            int         `json:"has_big_img"`
	IDBroadcast          int         `json:"id_broadcast"`
	PostsAnswerURI       string      `json:"posts__answer__uri"`
	PostsAnswerDate      string      `json:"posts__answer__date"`
	PostsAnswerText      string      `json:"posts__answer__text"`
	PostsAnswerClearText string      `json:"posts__answer__clear_text"` //amended
	PostsAnswerEmbeds    interface{} `json:"posts__answer__embeds"`     //amended
	PostsAnswerTime      string      `json:"posts__answer__time"`
	PostsAnswerType      int         `json:"posts__answer__type"`
	PostsAnswerAuthor    string      `json:"posts__answer__author"`
	PostsAnswerImages    interface{} `json:"posts__answer__images"`
	PostsAnswerAnswers   interface{} `json:"posts__answer__answers"`
	PostsAnswerTimestamp int         `json:"posts__answer__timestamp"`
}
type Posts struct {
	ID             int            `json:"id"`
	IDParent       interface{}    `json:"id_parent"`
	PostsURI       string         `json:"posts__uri"`
	HasBigImg      int            `json:"has_big_img"`
	PostsDate      string         `json:"posts__date"`
	PostsText      string         `json:"posts__text"`
	PostsClearText string         `json:"posts__clear_text"` //amended
	PostsEmbeds    interface{}    `json:"posts__embeds"`     //amended
	PostsTime      string         `json:"posts__time"`
	PostsType      int            `json:"posts__type"`
	IDBroadcast    int            `json:"id_broadcast"`
	PostsAuthor    string         `json:"posts__author"`
	PostsImages    []PostsImages  `json:"posts__images"`
	PostsAnswers   []PostsAnswers `json:"posts__answers"`
	PostsTimestamp int            `json:"posts__timestamp"`
}

package newspaper

import "time"

type Article struct {
	Author        string
	Title         string
	Text          string
	WrittenDate   time.Time
	PublishedDate time.Time
}

func NewArticle(author string, title string, text string) *Article {
	return &Article{
		Author:      author,
		Title:       title,
		Text:        text,
		WrittenDate: time.Now(),
	}
}

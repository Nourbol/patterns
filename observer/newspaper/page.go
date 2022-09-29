package newspaper

type Page struct {
	Articles []Article
}

func NewPage(articles ...Article) *Page {
	return &Page{Articles: articles}
}

func (p *Page) AddArticle(articles ...Article) {
	p.Articles = append(p.Articles, articles...)
}

func (p *Page) GetArticlesNumber() int {
	return len(p.Articles)
}

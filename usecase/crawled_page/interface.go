package crawled_page

import "github.com/d-jo/webcrawler/entity"

type Reader interface {
	GetChildren(url string) (map[string]int32, error)
	GetTree(url string) (*entity.CrawledPage, error)
}

type Writer interface {
	AddPage(page *entity.CrawledPage) error
}

type Repository interface {
	Reader
	Writer
}

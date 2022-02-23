package crawled_page

import "github.com/d-jo/webcrawler/entity"

type Reader interface {
	GetChildren(url string) (map[string]int32, error)
	GetTree(url string) (*entity.CrawledPage, error)
	GetAllKeys() ([]string, error)
}

type Writer interface {
	AddPage(page *entity.CrawledPage) error
	AddRootPage(url string) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetChildren(url string) (map[string]int32, error)
	GetTree(url string) (*entity.CrawledPage, error)
	AddPage(page *entity.CrawledPage) error
	GetAll() ([]*entity.CrawledPage, error)
	AddRootPage(url string) error
}

package crawled_page

import "github.com/d-jo/webcrawler/entity"

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service, error) {
	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) GetChildren(url string) (map[string]int32, error) {
	return s.repo.GetChildren(url)
}

func (s *Service) GetTree(url string) (*entity.CrawledPage, error) {
	return s.repo.GetTree(url)
}

func (s *Service) AddPage(page *entity.CrawledPage) error {
	return s.repo.AddPage(page)
}

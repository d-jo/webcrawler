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

func (s *Service) GetAll() ([]*entity.CrawledPage, error) {
	keys, err := s.repo.GetAllKeys()

	if err != nil {
		return nil, err
	}

	results := make([]*entity.CrawledPage, 0)

	for _, k := range keys {
		page, err := s.repo.GetTree(k)

		if err != nil {
			continue
		}

		results = append(results, page)
	}

	return results, nil
}

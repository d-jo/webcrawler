package crawler_mem

import (
	"fmt"
	"sync"

	"github.com/d-jo/webcrawler/entity"
)

/*
	The repository for storing pages. We use a map that maps to a map so that we get set-like behavior from it
	and we can count the occurances of a url. This means we dont have to remote duplicates when doing the regex
*/
type RepoService struct {
	Pages map[string]map[string]int32 `json:"pages"`
	mux   *sync.RWMutex
}

func NewRepoService() (*RepoService, error) {
	return &RepoService{
		Pages: make(map[string]map[string]int32),
		mux:   &sync.RWMutex{},
	}, nil
}

/*
	Gets the children for a url. Is not recursive, only returns the direct children
*/
func (s *RepoService) GetChildren(url string) (map[string]int32, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	pages, ok := s.Pages[url]

	var err error = nil
	if !ok {
		err = fmt.Errorf("url not found: %s", url)
	}

	return pages, err
}

/*
	Adds a crawled page to the repository
*/
func (s *RepoService) AddPage(page *entity.CrawledPage) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if existingPage, ok := s.Pages[page.Url]; ok {
		// exists, add to existing
		for i := range page.Children {
			existingPage[page.Children[i].GetUrl()] += 1
		}
	} else {
		// does not exist, allocate new map
		s.Pages[page.Url] = make(map[string]int32)
		for i := range page.Children {
			s.Pages[page.Url][page.Children[i].GetUrl()] += 1
		}
	}

	return nil
}

func (s *RepoService) getTreeRec(url string, visitedPages *map[string]bool) (*entity.CrawledPage, error) {
	if _, ok := (*visitedPages)[url]; ok {
		// we have been here before, dont get the children to avoid cycle
		return &entity.CrawledPage{Url: url}, nil
	}

	(*visitedPages)[url] = true

	// get the children
	children, err := s.GetChildren(url)

	if err != nil {
		// no children, key doesnt exist
		// oops
		// TODO
		return &entity.CrawledPage{Url: url}, nil
	}

	// we have the children

	page := entity.CrawledPage{
		Url:      url,
		Children: make([]*entity.CrawledPage, 0),
	}

	for child := range children {
		newChild, err := s.getTreeRec(child, visitedPages)

		if err != nil {
			continue
		}

		page.Children = append(page.Children, newChild)
	}

	return &page, nil
}

func (s *RepoService) GetTree(url string) (*entity.CrawledPage, error) {
	visitedPages := make(map[string]bool)

	root := entity.CrawledPage{
		Url:      url,
		Children: make([]*entity.CrawledPage, 0),
	}

	// for each child in root, create a crawled url and do the same for its children

	children := s.Pages[url]

	for child := range children {
		newChild, err := s.getTreeRec(child, &visitedPages)

		if err != nil {
			continue
		}

		root.Children = append(root.Children, newChild)
	}

	return &root, nil
}

func (s *RepoService) GetAllKeys() ([]string, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	var keys []string
	for k := range s.Pages {
		keys = append(keys, k)
	}

	return keys, nil
}

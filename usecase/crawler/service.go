package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/d-jo/webcrawler/entity"
	"github.com/d-jo/webcrawler/usecase/crawled_page"
	"github.com/d-jo/webcrawler/util"
)

type Service struct {
	pageService crawled_page.UseCase
	client      *http.Client
	workers     map[string]chan interface{}
	visited     map[string]bool
	mux         sync.RWMutex
}

func NewService(page_service crawled_page.UseCase) *Service {
	return &Service{
		pageService: page_service,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		workers: make(map[string]chan interface{}),
		visited: make(map[string]bool),
		mux:     sync.RWMutex{},
	}
}

func (s *Service) crawlWorker(url string, done <-chan interface{}) {

	go func() {
		currLink := url
		for {
			// check if we are done
			select {
			case <-done:
				return
			default:
			}

			// check if we have visited here before
			s.mux.RLock()
			visited, ok := s.visited[currLink]
			s.mux.RUnlock()

			if ok || visited {
				break
			}

			resp, err := s.client.Get(currLink)

			if err != nil {
				continue
			}

			if resp.StatusCode != http.StatusOK {
				continue
			}

			// check again, I dont want to do all the work
			// reading the body if we are done and need to exit
			select {
			case <-done:
				return
			default:
			}

			content, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				break
			}

			s.mux.Lock()
			s.visited[currLink] = true
			s.mux.Unlock()

			// use the util function for extracting links
			allLinks, err := util.SearchForURLs(string(content), url)

			// no links, this page doesnt lead anywhere
			// stop crawling this page
			if len(allLinks) == 0 {
				break
			}

			// set this workers next link
			currLink = allLinks[0][0]

			// for all the other links, start a goroutine
			for _, link := range allLinks[1:] {
				s.crawlWorker(link[0], done)
			}

			break
		}
	}()
}

func (s *Service) StartCrawling(sc *entity.StartCommand) *entity.GenericResponse {

	parsedUrl, err := url.Parse(sc.Url)

	if err != nil {
		return &entity.GenericResponse{
			Success: false,
			Message: "invalid url",
		}
	}

	done := make(chan interface{})
	host := parsedUrl.Hostname()

	// start going
	s.crawlWorker(host, done)

	s.workers[host] = done

	return &entity.GenericResponse{
		Success: true,
		Message: fmt.Sprintf("started crawling %s", host),
	}
}

func (s *Service) StopCrawling(sc *entity.StopCommand) *entity.GenericResponse {
	// TODO
	return &entity.GenericResponse{
		Success: false,
		Message: "not implemented",
	}
}

func (s *Service) List(lc *entity.ListCommand) *entity.ListResponse {

}

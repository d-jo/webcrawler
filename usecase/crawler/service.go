package crawler

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/d-jo/webcrawler/entity"
	"github.com/d-jo/webcrawler/usecase/crawled_page"
	"github.com/d-jo/webcrawler/util"
)

type Service struct {
	UnimplementedWebCrawlerServer
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

/*
	Crawl worker starts a goroutine that will crawl the given host. For each page,
	the URL and its children are sent to the repository for storage.

	The worker will start with the given URL. The first link it finds it will traverse next
	after it spawns goroutines to crawl all the other links

	we use the visited map to avoid crawling the same URL twice and running
	in to a cycle. If it fails to get a 200 response it will retry 5 times before finally
	giving up.

	there is no delay between requests. this should crawl the entire site web very quickly
	and terminate after the entire site has been crawled.
*/
func (s *Service) crawlWorker(target_url string, done <-chan interface{}) {
	go func() {
		currLink := target_url
		failures := 0
		for {
			log.Println("crawling", currLink)
			// check if we are done
			select {
			case <-done:
				return
			default:
			}

			if failures > 5 {
				return
			}

			// check if we have visited here before
			s.mux.RLock()
			visited, ok := s.visited[currLink]
			s.mux.RUnlock()

			// we have visited here before
			// so lets not make a request
			if ok || visited {
				break
			}

			// get the page
			resp, err := s.client.Get(currLink)

			if err != nil {
				failures++
				continue
			}

			if resp.StatusCode != http.StatusOK {
				failures++
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
				failures++
				continue
			}

			s.mux.Lock()
			s.visited[currLink] = true
			s.mux.Unlock()

			purl, _ := url.Parse(currLink)
			host := purl.Hostname()
			if len(purl.Port()) > 0 {
				host = fmt.Sprintf("%s:%s", host, purl.Port())
			}

			// use the util function for extracting links
			allLinks, err := util.SearchForURLs(string(content), host)

			if err != nil {
				// failed to extract links
				failures++
				continue
			}

			children := make([]*entity.CrawledPage, len(allLinks))
			for i, link := range allLinks {
				children[i] = &entity.CrawledPage{
					Url: link[0],
				}
			}

			// add to memory
			s.pageService.AddPage(&entity.CrawledPage{
				Url:      currLink,
				Children: children,
			})

			// no links, this page doesnt lead anywhere
			// stop crawling this page
			if len(allLinks) == 0 {
				break
			}

			// set this workers next link
			currLink = allLinks[0][0]

			if len(allLinks) > 1 {
				// for all the other links, start a goroutine
				for _, link := range allLinks[1:] {
					s.crawlWorker(link[0], done)
				}
			}
		}
	}()
}

func (s *Service) StartCrawling(ctx context.Context, sc *entity.StartCommand) (*entity.GenericResponse, error) {

	parsedUrl, err := url.Parse(sc.Url)

	if err != nil {
		return &entity.GenericResponse{
			Success: false,
			Message: "invalid url",
		}, nil
	}

	done := make(chan interface{})
	host := parsedUrl.Scheme + "://" + parsedUrl.Hostname()
	port := parsedUrl.Port()

	if len(port) > 0 {
		host = fmt.Sprintf("%s:%s", host, port)
	}

	log.Println("starting to crawl", host)

	s.pageService.AddRootPage(host)

	// start going
	s.crawlWorker(host, done)

	s.workers[host] = done

	return &entity.GenericResponse{
		Success: true,
		Message: fmt.Sprintf("started crawling %s", host),
	}, nil
}

func (s *Service) StopCrawling(ctx context.Context, sc *entity.StopCommand) (*entity.GenericResponse, error) {

	parsedUrl, err := url.Parse(sc.Url)

	if err != nil {
		return &entity.GenericResponse{
			Success: false,
			Message: "invalid url",
		}, nil
	}

	// close the channel
	close(s.workers[parsedUrl.Hostname()])

	return &entity.GenericResponse{
		Success: true,
		Message: fmt.Sprintf("stopped crawling %s", parsedUrl.Hostname()),
	}, nil
}

func (s *Service) List(ctx context.Context, lc *entity.ListCommand) (*entity.ListResponse, error) {

	pages, err := s.pageService.GetAll()

	if err != nil {
		return &entity.ListResponse{
			Success: false,
			Message: "failed to get pages",
		}, nil
	}

	return &entity.ListResponse{
		Success: true,
		Root:    pages,
	}, nil
}

func (s *Service) CloseAllWorkers() bool {
	for _, done := range s.workers {
		close(done)
	}
	return true
}

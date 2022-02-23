package crawler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/d-jo/webcrawler/entity"
	"github.com/d-jo/webcrawler/repository/crawler_mem"
	"github.com/d-jo/webcrawler/usecase/crawled_page"
	"github.com/d-jo/webcrawler/usecase/crawler"
	"github.com/d-jo/webcrawler/util"
)

func TestCrawl1(t *testing.T) {

	var handlerFunc http.HandlerFunc

	var pages map[string]string

	handlerFunc = func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("aksndkajsndkjasndkjasn"))
		p, ok := pages[r.URL.Path]
		if ok {
			w.Write([]byte(p))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}

	srv := httptest.NewServer(handlerFunc)
	defer srv.Close()

	pages = map[string]string{
		"/":      fmt.Sprintf("<a href=\"%s/page1\">page1</a>", srv.URL),
		"/page1": fmt.Sprintf("<a href=\"%s/page2\">page2</a>", srv.URL),
		"/page2": "empty page wiith no links",
	}

	memRepo, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Error(err)
	}

	pageService, err := crawled_page.NewService(memRepo)

	if err != nil {
		t.Error(err)
	}

	crawler := crawler.NewService(pageService)
	defer crawler.CloseAllWorkers()

	sc := entity.StartCommand{
		Url: srv.URL + "/",
	}

	resp := crawler.StartCrawling(&sc)

	if resp.Success != true {
		t.Error("expected success:", resp.Message)
	}

	lc := entity.ListCommand{}

	time.Sleep(time.Second * 2)

	list := crawler.List(&lc)

	var buf strings.Builder
	util.FPrintChildren(&buf, list.Root[0], 0)

	t.Logf("%+v", len(list.Root))
	t.Log(buf.String())

	if len(list.Root) != 1 {
		t.Error("expected 1 root page")
	}

	if len(list.Root[0].Children) != 1 {
		t.Error("expected 1 child page")
	}

	if len(list.Root[0].Children[0].Children) != 1 {
		t.Error("expected 0 child pages")
	}

	if len(list.Root[0].Children[0].Children[0].Children) != 0 {
		t.Error("expected 0 child pages")
	}
}

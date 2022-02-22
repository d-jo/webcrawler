package crawler_mem_test

import (
	"testing"

	"github.com/d-jo/webcrawler/entity"
	"github.com/d-jo/webcrawler/repository/crawler_mem"
)

func TestAddPageCold(t *testing.T) {
	repoService, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Errorf("error creating repo service: %s", err)
	}

	// list of child urls to add
	expected := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/1",
		},
		{
			Url: "http://www.example.com/2",
		},
	}
	expectedLen := len(expected)

	// create the page obj
	page := entity.CrawledPage{
		Url:      "http://www.example.com",
		Children: expected,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	pages, err := repoService.GetChildren("http://www.example.com")

	if err != nil {
		t.Errorf("error getting children: %s", err)
	}

	// check to see if a copy or references is stored in the map
	if len(pages) != expectedLen {
		t.Errorf("expected %d children, got %d", expectedLen, len(pages))
	}

	// check the pages
	for i := range expected {
		_, ok := pages[expected[i].GetUrl()]

		if !ok {
			t.Errorf("expected map to have %s, got %v", expected[i], ok)
			continue
		}
	}
}

func TestAddPageHot(t *testing.T) {
	repoService, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Errorf("error creating repo service: %s", err)
	}

	// list of child urls to add
	expected := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/1",
		},
		{
			Url: "http://www.example.com/2",
		},
	}
	expectedLen := len(expected)

	// create the page obj
	page := entity.CrawledPage{
		Url:      "http://www.example.com",
		Children: expected,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	expected2 := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/3",
		},
		{
			Url: "http://www.example.com/4",
		},
	}
	expectedLen += len(expected2)

	page = entity.CrawledPage{
		Url:      "http://www.example.com",
		Children: expected2,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	pages, err := repoService.GetChildren("http://www.example.com")

	if err != nil {
		t.Errorf("error getting children: %s", err)
	}

	// check to see if a copy or references is stored in the map
	if len(pages) != expectedLen {
		t.Errorf("expected %d children, got %d", expectedLen, len(pages))
	}

	// check the pages
	for i := range expected {
		_, ok := pages[expected[i].GetUrl()]
		_, ok2 := pages[expected2[i].GetUrl()]

		if !ok && !ok2 {
			t.Errorf("expected map to have %s, got %v and `%v", expected[i], ok, ok2)
			continue
		}
	}
}

func TestGetPagesNotExisting(t *testing.T) {
	repoService, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Errorf("error creating repo service: %s", err)
	}

	_, err = repoService.GetChildren("http://www.example.com")

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestDuplicateChildren(t *testing.T) {
	repoService, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Errorf("error creating repo service: %s", err)
	}

	// list of child urls to add
	expected := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/1",
		},
		{
			Url: "http://www.example.com/1",
		},
	}
	expectedLen := 1

	// create the page obj
	page := entity.CrawledPage{
		Url:      "http://www.example.com",
		Children: expected,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	pages, err := repoService.GetChildren("http://www.example.com")

	t.Logf("%+v", pages)

	if err != nil {
		t.Errorf("error getting children: %s", err)
	}

	// check to see if a copy or references is stored in the map
	if len(pages) != expectedLen {
		t.Errorf("expected %d children, got %d", expectedLen, len(pages))
	}

	// check the pages
	for i := range expected {
		v, ok := pages[expected[i].GetUrl()]

		if !ok {
			t.Errorf("expected map to have %s, got %v", expected[i], ok)
		}

		if v != 2 {
			t.Errorf("expected map to have %s, got %v", expected[i], v)
		}
	}
}

func TestGetTreeBasic(t *testing.T) {
	repoService, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Errorf("error creating repo service: %s", err)
	}

	// list of child urls to add
	expected := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/1",
		},
		{
			Url: "http://www.example.com/2",
		},
	}

	// create the page obj
	page := entity.CrawledPage{
		Url:      "http://www.example.com",
		Children: expected,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	// list of child urls to add
	expected2 := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/3",
		},
		{
			Url: "http://www.example.com/4",
		},
	}

	// create the page obj
	page2 := entity.CrawledPage{
		Url:      "http://www.example.com/1",
		Children: expected2,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page2)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	// so now here it should be root leads to 1 and 2
	// and 1 leads to 3 and 4

	// so lets do the tree

	tree, err := repoService.GetTree("http://www.example.com")

	if err != nil {
		t.Errorf("error getting tree: %s", err)
	}

	// check the tree

	topLevelChildren := tree.GetChildren()

	if len(topLevelChildren) != 2 {
		t.Errorf("expected 2 children, got %d", len(topLevelChildren))
		t.Logf("%+v", topLevelChildren)
	}

	t.Logf("root: %+v (%+v)", tree.GetUrl(), tree.GetChildren())

	for _, ch := range topLevelChildren {
		t.Logf("ch: %+v (%+v)", ch.GetUrl(), ch.GetChildren())
		if ch.GetUrl() == "http://www.example.com/1" {
			oneChildren := len(ch.GetChildren())
			if oneChildren != 2 {
				t.Errorf("expected 2 children for example..1 , got %d", oneChildren)
			}
		} else if ch.GetUrl() == "http://www.example.com/2" {
			twoChildren := len(ch.GetChildren())
			if twoChildren != 0 {
				t.Errorf("expected 0 children for example..2 , got %d", twoChildren)
			}
		}

	}

}

/*
// behaves as expected but I havent written the checking code yet so comment out for this commit
func TestGetTreeCycle(t *testing.T) {
	repoService, err := crawler_mem.NewRepoService()

	if err != nil {
		t.Errorf("error creating repo service: %s", err)
	}

	// list of child urls to add
	expected := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/1",
		},
		{
			Url: "http://www.example.com/2",
		},
	}

	// create the page obj
	page := entity.CrawledPage{
		Url:      "http://www.example.com",
		Children: expected,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	// list of child urls to add
	expected2 := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/2",
		},
	}

	// create the page obj
	page2 := entity.CrawledPage{
		Url:      "http://www.example.com/1",
		Children: expected2,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page2)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	// list of child urls to add
	expected3 := []*entity.CrawledPage{
		{
			Url: "http://www.example.com/1",
		},
	}

	// create the page obj
	page3 := entity.CrawledPage{
		Url:      "http://www.example.com/2",
		Children: expected3,
	}

	// add the pages to the memory map
	err = repoService.AddPage(&page3)

	if err != nil {
		t.Errorf("error adding page: %s", err)
	}

	// so now here it should be root leads to 1 and 2
	// and 1 leads to 3 and 4

	// so lets do the tree

	tree, err := repoService.GetTree("http://www.example.com")

	if err != nil {
		t.Errorf("error getting tree: %s", err)
	}

	// check the tree

	topLevelChildren := tree.GetChildren()

	if len(topLevelChildren) != 2 {
		t.Errorf("expected 2 children, got %d", len(topLevelChildren))
		t.Logf("%+v", topLevelChildren)
	}

	t.Logf("root: %+v (%+v)", tree.GetUrl(), tree.GetChildren())

	for _, ch := range topLevelChildren {
		t.Logf("ch: %+v (%+v)", ch.GetUrl(), ch.GetChildren())
	}

	t.Fail()

}
*/

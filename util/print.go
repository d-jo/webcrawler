package util

import (
	"fmt"
	"io"
	"strings"

	"github.com/d-jo/webcrawler/entity"
)

func FPrintChildren(w io.Writer, p *entity.CrawledPage, depth int) {
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth), p.Url)
	for _, child := range p.Children {
		FPrintChildren(w, child, depth+1)
	}
}

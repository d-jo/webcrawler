package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/d-jo/webcrawler/entity"
	cb "github.com/d-jo/webcrawler/usecase/crawler"
	"github.com/d-jo/webcrawler/util"
	"google.golang.org/grpc"
)

var (
	startValue string
	stopValue  string
	listValue  bool
)

func main() {

	flag.StringVar(&startValue, "start", "", "start crawling")
	flag.StringVar(&stopValue, "stop", "", "stop crawling")
	flag.BoolVar(&listValue, "list", false, "list all crawled pages")

	flag.Parse()

	conn, err := grpc.Dial("localhost:8985", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	client := cb.NewWebCrawlerClient(conn)

	if startValue != "" {
		log.Println("starting")
		resp, err := client.StartCrawling(context.Background(), &entity.StartCommand{Url: startValue})

		if err != nil {
			panic(err)
		}

		if resp.GetSuccess() {
			log.Println("started crawling ", startValue)
		} else {
			log.Println("failed to start crawling:", resp.GetMessage())
		}
	}

	if stopValue != "" {
		log.Println("stopping")
		resp, err := client.StopCrawling(context.Background(), &entity.StopCommand{Url: stopValue})

		if err != nil {
			panic(err)
		}

		if resp.GetSuccess() {
			log.Println("stopped crawling ", stopValue)
		} else {
			log.Println("failed to stop crawling:", resp.GetMessage())
		}
	}

	if listValue {
		resp, err := client.List(context.Background(), &entity.ListCommand{})

		if err != nil {
			panic(err)
		}

		if resp.GetSuccess() {
			log.Println("crawled pages:")

			for _, root := range resp.GetRoot() {
				var buf strings.Builder

				util.FPrintChildren(&buf, root, 0)

				fmt.Println(buf.String())
			}
		} else {
			log.Println("failed to list all crawled pages:", resp.GetMessage())
		}
	}

}

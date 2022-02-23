package main

import (
	"log"
	"net"

	"github.com/d-jo/webcrawler/repository/crawler_mem"
	"github.com/d-jo/webcrawler/usecase/crawled_page"

	cb "github.com/d-jo/webcrawler/usecase/crawler"
	"google.golang.org/grpc"
)

func main() {

	memRepo, err := crawler_mem.NewRepoService()

	if err != nil {
		panic(err)
	}

	pageService, err := crawled_page.NewService(memRepo)

	if err != nil {
		panic(err)
	}

	crawler := cb.NewService(pageService)
	defer crawler.CloseAllWorkers()

	grpcServer := grpc.NewServer()

	cb.RegisterWebCrawlerServer(grpcServer, crawler)

	log.Println("starting gRPC on port :8985")
	lis, err := net.Listen("tcp", ":8985")

	if err != nil {
		panic(err)
	}

	grpcServer.Serve(lis)
}

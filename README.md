# gRPC Crawler

## Client

`crawl --start <url>` - starts crawling the given url

`crawl --stop <url>` - stops crawling the given url

`crawl --list` - show the site tree


## Service

The service should crawl pages on the same domain only. The service should be able to handle cycles and execute requests concurrently.

## Building

The `Makefile` contains the following targets:

`all` - builds the client, testserver, and server and all their dependencies

`server` - builds the server and its target, including the gRPC targets

`testserver` - builds the test server and its target, including the gRPC targets

`client` - builds the client and its target, including the gRPC targets

`grpc_entity` - builds the gRPC entities, included in the `entity/entity.proto` file

`grpc_service` - builds the gRPC service, included in the `usecase/crawler/crawler.proto` file. Depends on the `grpc_entity` target

`grpc_clean` - cleans the generated gRPC files for entity and service

`clean` - cleans compiled files from release folder and from their folders in api and cmd

## Source Map

`api`
  -  `server` - the gRPC server
  -  `testserver` - a simple test http server to test the crawler locally

`cmd`
  - `client` - the gRPC client

`entity`
  - `entity.proto` - the entity definitions

`repository `
  - `crawler_mem` 
    - `crawler_mem.go` - the repository implementation for storing visited sites in memory
    - `crawler_mem_test.go` - the repository test file

`usecase`
  - `crawled_page`
    - `interface.go` - the interface for the crawled page usecase and repository
    - `service.go` - the service implementation for the crawled page usecase

  - `crawler`
    - `service.go` - the usecase implementation
    - `service_test.go` - the usecase test file
    - `crawler.proto` - the usecase service proto definition
    - `crawler_grpc.pb.go` - generated gRPC file
    - `crawler.pb.go` - generated gRPC file

`util`
  - `print.go` - the print utility for printing the tree
  - `url.go` - the url utility for parsing urls and searching for links

`Makefile` - the makefile for building the project. see above for possible targets

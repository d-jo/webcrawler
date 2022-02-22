# gRPC Crawler

## Client

`crawl -start <url>` - starts crawling the given url

`crawl -stop <url>` - stops crawling the given url

`crawl -list` - show the site tree


## Service

The service should crawl pages on the same domain only. The service should be able to handle cycles and execute requests concurrently.

## Building

The `Makefile` contains the following targets:

`all` - builds the client and server and all their dependencies

`server` - builds the server and its target, including the gRPC targets

`client` - builds the client and its target, including the gRPC targets

`grpc_entity` - builds the gRPC entities, included in the `entity/entity.proto` file

`grpc_service` - builds the gRPC service, included in the `usecase/crawler/crawler.proto` file. Depends on the `grpc_entity` target

`grpc_clean` - cleans the generated gRPC files for entity and service

`clean` - cleans compiled files from release folder and from their folders in api and cmd


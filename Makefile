
CLIENTNAME = crawl
SERVERNAME = server

all: client server

server: ./api/server/main.go grpc_service 
	mkdir -p release
	go build --tags=prod -o ./release/$(SERVERNAME) ./api/server/main.go && cp ./release/$(SERVERNAME) ./api/server/$(SERVERNAME)

client: ./cmd/client/main.go grpc_service
	mkdir -p release
	go build --tags=prod -o ./release/$(CLIENTNAME)./cmd/client/main.go && cp ./release/$(CLIENTNAME) ./cmd/client/$(CLIENTNAME)

grpc_entity: ./entity/entity.proto 
	protoc -I=./entity --go_opt=paths=source_relative --go_out=./entity/ --go-grpc_out=./entity/ --go-grpc_opt=paths=source_relative entity/*.proto

grpc_service: ./usecase/crawler/crawler.proto grpc_entity
	protoc --go_opt=paths=source_relative --go_out=./usecase/crawler --go-grpc_opt=paths=source_relative --go-grpc_out=./usecase/crawler --proto_path=./entity --proto_path=./usecase/crawler usecase/crawler/crawler.proto

grpc_clean:
	rm -f ./entity/entity.pb.go
	rm -f ./usecase/crawler/crawler.pb.go

clean: grpc_clean
	rm -f ./release/$(SERVERNAME)
	rm -f ./api/server/$(SERVERNAME)
	rm -f ./release/$(CLIENTNAME)
	rm -f ./cmd/client/$(CLIENTNAME)

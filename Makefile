postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb: 
	sudo docker exec -it postgres createdb --username=root --owner=root register
	
dropdb:
	sudo docker exec -it postgres dropdb  students

migrateup: 
	migrate -path migration/ -database "postgresql://root:secret@localhost:5435/register?sslmode=disable" -verbose up

migratedown:
	migrate -path migration/ -database "postgresql://root:secret@localhost:5435/register?sslmode=disable" -verbose down

proto:
	protoc --go_out=./pkg/admin --go_opt=paths=source_relative \
    --go-grpc_out=./api/proto --go-grpc_opt=paths=source_relative \
    api/proto/service.proto



.PHONY: postgres createdb migrateup run
postgres:
	docker run --name postgres14.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres14.2 createdb --username=root --owner=root venta-de-campos-backend

dropdb:
	docker exec -it postgres14.2 dropdb venta-de-campos-backend

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/venta-de-campos-backend?sslmode=disable" -verbose up	

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/venta-de-campos-backend?sslmode=disable" -verbose down	

test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown test


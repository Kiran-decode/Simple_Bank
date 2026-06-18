postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root1 -e POSTGRES_PASSWORD=ananya -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root1 --owner=root1 SimpleBank

dropdb:
	docker exec -it postgres13 dropdb SimpleBank

migrateup:
	migrate -source file://$(PWD)/db/migration \
	-database "postgresql://root1:ananya@localhost:5432/SimpleBank?sslmode=disable" \
	up

migratedown:
	migrate -source file://$(PWD)/db/migration \
	-database "postgresql://root1:ananya@localhost:5432/SimpleBank?sslmode=disable" \
	down
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
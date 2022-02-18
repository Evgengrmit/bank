postgres:
	docker run --name postgres14 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d postgres:latest
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root bank
dropdb:
	docker exec -it postgres14 dropdb bank
migrateup:
	migrate -path db/migration -database "postgresql://root:qwerty@localhost:5432/bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:qwerty@localhost:5432/bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
install:
	go mod download
up:
	docker-compose up -d

upv:
	docker-compose up

down:
	docker-compose down

new_migration:
	migrate create -ext sql -dir migration -seq $(name)

migrateup:
	migrate -path migration -database "$(DB_SOURCE)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_SOURCE)" -verbose down 1

server:
	go run main.go

.PHONY: install up upv down new_migration migrateup migratedown migrateup1 migratedown1 server

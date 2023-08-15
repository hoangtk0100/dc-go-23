include .env
export

network:
	docker network create $(NETWORK_NAME)

postgres:
	docker run --name $(DB_CONTAINER_NAME) --network $(NETWORK_NAME) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d $(PSQL_IMAGE)

createdb:
	docker exec -it $(DB_CONTAINER_NAME) createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) $(DB_NAME)

dropdb:
	docker exec -it $(DB_CONTAINER_NAME) dropdb $(DB_NAME)

db:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_USERNAME) $(DB_NAME)

server:
	go run main.go

.PHONY: network postgres createdb dropdb db server

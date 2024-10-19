run-instance: cmd/server/main.go
	go run cmd/server/main.go

run-distributed: cmd/server/main.go
	docker-compose up --build

down-distributed:
	docker-compose down

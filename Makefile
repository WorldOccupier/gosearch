.PHONY: up down rebuild logs clean localbuild lint fmt

up:
	docker compose up -d

down:
	docker compose down

rebuild:
	docker compose down -v && docker compose up -d --build && docker compose logs -f

logs:
	docker compose logs -f

clean:
	docker compose down -v

lint:
	cd cmd/main && golangci-lint run ./...

fmt:
	cd cmd/main && go fmt ./...

localbuild:
	go build -C cmd/main -o ../../gosearch .

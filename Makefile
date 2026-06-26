up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

build:
	cd order-service && go build -o app cmd/main.go

run:
	cd order-service && go run cmd/main.go
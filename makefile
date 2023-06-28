swagger-docs:
	swag init --parseDependency --parseInternal --parseDepth 1 -d cmd/api

run:
	go run cmd/api/main.go

gen:
	go run cmd/generate/main.go

api:
	go run cmd/api/main.go

up:
	go run cmd/migrate/main.go -up

up_test:
	go run cmd/migrate/main.go -up -test

down:
	go run cmd/migrate/main.go -down

down_test:
	go run cmd/migrate/main.go -down -test

force:
	go run cmd/migrate/main.go -force -version 1

swagger-docs-generate:
	swag init --parseDependency --parseInternal --parseDepth 1 -d cmd/api

# docker exec -it <container-id> sh
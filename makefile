swagger-docs:
	swag init --parseDependency --parseInternal --parseDepth 1 -d cmd/api

run:
	go run cmd/api/main.go

gen:
	go run cmd/generate/main.go

api:
	go run cmd/api/main.go

# docker exec -it <container-id> sh
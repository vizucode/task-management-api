.PHONY: swagger-install swagger-gen swagger-serve run

# Install swagger dependencies
swagger-install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/fiber-swagger
	go get -u github.com/swaggo/files

# Generate swagger documentation
swagger-gen:
	swag init -g main.go --output ./docs

# Serve swagger documentation (after running the app)
swagger-serve:
	@echo "Swagger UI available at: http://localhost:8081/swagger/index.html"

# Run the application
run:
	go run main.go

# Generate swagger and run
dev: swagger-gen run

# Clean generated docs
clean:
	rm -rf ./docs

# Install dependencies and generate swagger
setup: swagger-install swagger-gen
.PHONY: default run build test docs clean

APP_NAME=MyFlix


default: run-with-docs

run:
	# Clear console on start-up
	@clear
	# Starting Api
	@echo "Starting Api In Debug Mode..."
	@go run main.go
run-with-docs:
	# Clear console on start-up
	@clear
	# Starting Swagger
	@echo "Starting Swagger..."
	@swag init
	# Starting Api
	@echo "Starting Api..."
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rm ./docs
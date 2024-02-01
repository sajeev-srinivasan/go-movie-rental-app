GO = go
TARGET = movie-rental-app
APP_EXECUTABLE = "./target/$(TARGET)"
ENV = "local"
CONFIG_FILE_PATH = "./setup/env/$(ENV).yaml"
HTTP_SERVE = "http-serve"
MIGRATE = "migrate"

build:
	$(GO) build -o $(APP_EXECUTABLE) cmd/main.go

run: build
	$(APP_EXECUTABLE)

test:
	$(GO) test ./...

clean:
	rm -f ./target/$(TARGET)

http-serve: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE_PATH) $(HTTP_SERVE)

migrate: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE_PATH) $(MIGRATE)
GO = go
TARGET = movie-rental-app

build:
	$(GO) build -o ./target/$(TARGET) cmd/main.go

run: build
	./target/$(TARGET)

test:
	$(GO) test ./...

clean:
	rm -f ./target/$(TARGET)
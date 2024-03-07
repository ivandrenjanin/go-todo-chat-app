.PHONY: clean build run dev sqlc templ migrate test

BUILD_DIR = ./bin

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

build:
	go build -o $(BUILD_DIR)/ ./main.go

run: sqlc tailwind templ build 
	$(BUILD_DIR)/main serve ${ARGS}

test:
	go test -v ./test/...

migrate: build
	$(BUILD_DIR)/main migrate:up ${ARGS}

playground: build
	$(BUILD_DIR)/main pg ${ARGS}

dev: 
	air -c .air.toml

sqlc:
	sqlc generate

templ:
	templ generate

tailwind:
	npm run tailwind

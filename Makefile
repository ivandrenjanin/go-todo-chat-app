.PHONY: clean build run dev sqlc templ migrate dcu dcd

APP_NAME = go-chat-app
BUILD_DIR = ./bin

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

build:
	go build -o $(BUILD_DIR)/ ./cmd/***

run: sqlc tailwind templ build 
	$(BUILD_DIR)/server ${ARGS}

migrate: build
	$(BUILD_DIR)/migrations ${ARGS}

dev: 
	air -c .air.toml

sqlc:
	sqlc generate

templ:
	templ generate

tailwind:
	npm run tailwind

dcu:
	docker-compose up --detach

dcd:
	docker-compose down

.PHONY: clean build run dev sqlc templ migrate dcu dcd

APP_NAME = go-chat-app
BUILD_DIR = ./bin

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/server/main.go

run: build 
	$(BUILD_DIR)/$(APP_NAME) ${ARGS}

dev: 
	air -c .air.toml

sqlc:
	sqlc generate

templ:
	templ generate

migrate:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/migrations/main.go && $(BUILD_DIR)/$(APP_NAME) 


dcu:
	docker-compose up --detach

dcd:
	docker-compose down

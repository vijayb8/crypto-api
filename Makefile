PROJECT_NAME := crypto-api
PROJECT := github.com/vijayb8/$(PROJECT_NAME)
PKG := $(PROJECT)/cmd/$(PROJECT_NAME)
PKG_LIST := $(shell go list ${PROJECT}/... | grep -v /vendor/)
BUILD_DIR := build
LDFLAGS := -ldflags "-X main.Version=`git rev-parse HEAD`"

.PHONY: all test build clean help

all: help

build:
	go build -o ${BUILD_DIR}/${PROJECT_NAME} ${LDFLAGS} ${PKG}
	
install: ## Run app in local environment
	docker-compose up --build api

stop: ## Stop all app containers in local environment
	docker-compose down

docker-clean: ## Stop running app in local environment and remove all images and volumes for it
	docker-compose down --rmi local --volumes --remove-orphans

clean: ## Remove generated files after build
	rm -rf ${BUILD_DIR}

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
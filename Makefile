PROJECT_NAME := crypto-api
PROJECT := github.com/vijayb8/$(PROJECT_NAME)
PKG := $(PROJECT)/cmd/$(PROJECT_NAME)
PKG_LIST := $(shell go list ${PROJECT}/... | grep -v /vendor/)
BUILD_DIR := build
LDFLAGS := -ldflags "-X main.Version=`git rev-parse HEAD`"

.PHONY: all dep lint test build clean help

all: help

dep: ## Get the dependencies
	dep version || go get -u github.com/golang/dep/cmd/dep # install dep if not exist
	dep ensure

build: dep ## Build the binary file
	packr help || go get -u github.com/gobuffalo/packr/...  # install packr if not exist
	packr
	go build -o ${BUILD_DIR}/${PROJECT_NAME} ${LDFLAGS} ${PKG}
	packr clean

before-commit: lint test swagger ## Run checks and update swagger info

dev-up: ## Run app in local environment
	# docker-compose up -d
	docker-compose up --build api

dev-down: ## Stop all app containers in local environment
	docker-compose down

dev-clean: ## Stop running app in local environment and remove all images and volumes for it
	docker-compose down --rmi local --volumes --remove-orphans

clean: ## Remove generated files after build
	rm -rf ${BUILD_DIR}

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
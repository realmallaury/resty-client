mod: ## run go mod
	GO111MODULE=on go mod tidy

up: ## start docker containers
	docker-compose up -d

down: ## stop docker containers
	docker-compose stop && docker-compose rm -f && docker system prune -f

test: ## Build project as local docker image for running tests and start the containers 
	docker-compose up --build --abort-on-container-exit && docker-compose stop && docker-compose rm -f && docker system prune -f

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
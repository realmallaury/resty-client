mod:
	GO111MODULE=on go mod tidy

up:
	docker-compose up -d

down:
	docker-compose stop && docker-compose rm -f && docker system prune -f

test:
	docker-compose up --build --abort-on-container-exit && docker-compose rm -f && docker system prune -f
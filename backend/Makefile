.PHONY: test
test:
	docker-compose up -d
	go run ./migrate/. up
	go clean -testcache
	ENV=test DYNAMO_ENDPOINT=http://localhost:8000 go test -v ./app/...

.PHONY: up
up:
	docker-compose up -d
	go run ./migrate/. up

.PHONY: down
down:
	docker-compose down

.PHONY: dev
dev:
	make up
	sh ./dev.sh

.PHONY: deploy
deploy:
	sh ./deploy.sh


.PHONY: test
test:
	docker-compose up -d
	go run ./migrate/. up
	go clean -testcache
	ENV=test DYNAMO_ENDPOINT=http://localhost:8000 go test -v ./app/...

.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: migrate-up
migrate-up:
	go run ./migrate/. up

.PHONY: migrate-down
migrate-down:
	go run ./migrate/. down

.PHONY: dev
dev:
	make docker-up
	sh ./dev.sh

.PHONY: deploy
deploy:
	sh ./deploy.sh


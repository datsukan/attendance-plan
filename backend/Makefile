.PHONY: test
test:
	docker-compose up -d
	go run ./migrate/. up
	export MODE_UT=TRUE
	go clean -testcache
	export ENV=test
	go test -v ./app/...

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


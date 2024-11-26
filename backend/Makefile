.PHONY: test
test:
	docker-compose up -d; \
	export MODE_UT=TRUE; \
	go clean -testcache; \
	export Env=test; \
	go test -v ./domain/...

.PHONY: up
up:
	docker-compose up -d
	go run ./migrate/. up

.PHONY: down
down:
	docker-compose down

.PHONY: dev
dev:
	sam build; \
	sam local start-api --env-vars env.json --docker-network backend_default

.PHONY: deploy
deploy:
	sam build; \
	sam deploy --parameter-overrides $(cat .prd.env)


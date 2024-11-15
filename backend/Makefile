.PHONY: test-domain
test-domain:
	docker compose up -d; \
	export MODE_UT=TRUE; \
	go clean -testcache; \
	go test -v -coverprofile=./c.out ./domain/...; \
	# go tool cover -func=c.out; \
	docker compose down; \
	rm ./c.out

.PHONY: test-domain
test-domain:
	docker compose up -d; \
	export MODE_UT=TRUE; \
	go clean -testcache; \
	go test -v ./domain/...; \
	#docker compose down; \

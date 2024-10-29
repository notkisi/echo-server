.PHONY: get
get:
	curl "http://localhost:8080/echo?message=test%20get%20message"
	@echo

.PHONY: post
post:
	curl -X POST http://localhost:8080/echo -H "Content-Type: application/json" -d '{"message": "test post message"}'
	@echo

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v -short ./...
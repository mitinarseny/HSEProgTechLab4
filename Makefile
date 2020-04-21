.PHONY: test
test:
	go test -v -bench ./...

.PHONY: report
report:
	docker-compose up

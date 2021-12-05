release:
	@> version.go; \
	echo "package gomodel" >> version.go; \
	echo "" >> version.go; \
	echo "const VERSION = \"$(version)\"" >> version.go; \
	git add ./version.go; \
	git commit -m "chore(release): $(version)"; \
	git tag $(version); \
	git push; \
	git push --tags; \
	go list -m github.com/metauro/gomodel@$(version)

install:
	go install ./cmd/gomodel

build:
	go build -o gomodel ./cmd/gomodel/main.go

.PHONY: test
test:
	go run ./cmd/gomodel/main.go gen  -o test --dsn "root:root@(localhost:3306)/dev" --table gomodel
	go test ./test

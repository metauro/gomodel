release:
	@> version.go; \
	echo "package gomodel" >> version.go; \
	echo "" >> version.go; \
	echo "const VERSION = \"$(version)\"" >> version.go; \
	git add ./version.go; \
	git commit -m "chore(release): $(version)"; \
	git tag $(version); \
	git push; \
	git push --tags \
	go list -m github.com/metauro/gomodel@$(version)

install:
	go install ./cmd/gomodel

build:
	go build -o gomodel ./cmd/gomodel/main.go

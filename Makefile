release:
	@> version.go; \
	echo "package gomodel" >> version.go; \
	echo "" >> version.go; \
	echo "const VERSION = \"$$(git describe --abbrev=0 --tags)\"" >> version.go; \
	git add .; \
	git commit -m "chore(release): $$(git describe --abbrev=0 --tags)"; \
	git push; \
	git push --tags


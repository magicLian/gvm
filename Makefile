.PHONY: gomod release clean

# Download go modules
gomod:
	go mod download
	go mod vendor

release: gomod
	goreleaser release --clean

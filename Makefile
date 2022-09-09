build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist

tidy:
	go fmt ./...
	go mod tidy

release:
	goreleaser release --rm-dist

test:
	go test
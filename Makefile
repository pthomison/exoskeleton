build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist

tidy:
	go fmt ./...
	go mod tidy

test:
	go test ./... -v

delve-test:
	dlv test ./cmd


NEXT_TAG=$(shell exoskeleton rev -i $(shell git tag --list | tail -n 1))
release:
	git tag $(NEXT_TAG)
	git push origin $(NEXT_TAG)


# dog food
generate-readme:
	go run ./... template -i ./TEMPLATE.md -o ./README.md 
EXOSKELETON=go run ./...

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

# dog food
NEXT_TAG=$(shell $(EXOSKELETON) rev -i $(shell git tag --list --sort version:refname | tail -n 1))
release:
	git tag $(NEXT_TAG)
	git push origin $(NEXT_TAG)

generate-readme:
	@$(EXOSKELETON) template -i ./TEMPLATE.md -o ./README.md \
	-v 'ROOT_HELP=$(shell $(GR) --help | base64)' \
	-v 'TEMPLATE_HELP=$(shell $(GR) template --help | base64)' \
	-v 'INJECTOR_HELP=$(shell $(GR) ssm-k8s-injector --help | base64)' \
	-v 'REV_HELP=$(shell $(GR) rev --help | base64)' \
	-v 'ETHPRICE_HELP=$(shell $(GR) ethprice --help | base64)'

	git add ./README.md && git commit -m "README generation" || true

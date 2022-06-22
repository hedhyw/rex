GOLANG_CI_LINT_VER:=v1.46.2
COVER_PACKAGES=${shell go list ./... | grep -Ev 'test' | tr '\n' ','}

all: lint test
.PHONY: all

lint: bin/golangci-lint
	./bin/golangci-lint run
.PHONY: lint

test:
	go test \
		-coverpkg=${COVER_PACKAGES} \
		-covermode=count \
		-coverprofile=coverage.out \
		./...
	go tool cover -func=coverage.out
.PHONY: test

test.fuzz:
	# make test.fuzz NAME=FuzzIPv4
	# make test.fuzz NAME=FuzzRangeNUmber
	go test -fuzz $(NAME) "github.com/hedhyw/rex/pkg/dialect/base"
.PHONY: test.fuzz

tidy:
	go mod tidy
.PHONY: vendor

bin/golangci-lint:
	curl \
		-sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s $(GOLANG_CI_LINT_VER)

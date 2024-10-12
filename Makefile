GOLANG_CI_LINT_VER:=v1.61.0
COVER_PACKAGES=${shell go list ./... | grep -Ev 'test|cmd' | tr '\n' ','}

all: lint test
.PHONY: all

lint: bin/golangci-lint-$(GOLANG_CI_LINT_VER)
	./bin/golangci-lint-$(GOLANG_CI_LINT_VER) run
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
	# make test.fuzz NAME=FuzzRangeNumber
	go test -fuzz $(NAME) "github.com/hedhyw/rex/pkg/dialect/base"
.PHONY: test.fuzz

tidy:
	go mod tidy
.PHONY: vendor

bin/golangci-lint-$(GOLANG_CI_LINT_VER):
	curl \
		-sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s $(GOLANG_CI_LINT_VER)
	mv bin/golangci-lint bin/golangci-lint-$(GOLANG_CI_LINT_VER)

build:
	go build -o ./bin/rex ./cmd/generator/main.go

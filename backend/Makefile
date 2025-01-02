## Colors
COLOR_RESET   = \033[0m
COLOR_INFO    = \033[32m
COLOR_COMMENT = \033[33m

main_package_path = ./cmd/app
binary_name = app

.PHONY: help
## Help
help:
	@printf "${COLOR_COMMENT}Usage:${COLOR_RESET}\n"
	@printf " make [target] [args...]\n\n"
	@printf "${COLOR_COMMENT}Available targets:${COLOR_RESET}\n"
	@awk '/^[a-zA-Z\-\0-9\.@]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf " ${COLOR_INFO}%-16s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: tidy
## format code and tidy modfile
tidy:
	go mod tidy -v
	go fmt ./...


.PHONY: audit
## run quality control checks
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


.PHONY: test
## run unit tests
test:
	go test -race -buildvcs -vet=off ./...


.PHONY: build
## build the application
build:
	go build -o=./build/${binary_name} ${main_package_path}

.PHONY: run
## run the  application
run:
	./build/${binary_name}

.PHONY: run-dev
## run the application with reloading on file changes
run-dev:
	go run github.com/air-verse/air@latest \
		--build.cmd "make build" --build.bin "./build/${binary_name}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, sql" \
		--misc.clean_on_exit "true"

.PHONY: migrations/new
## create a new database migration
migrations/new:
	@go run -tags 'postgresql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=./migrations ${name}
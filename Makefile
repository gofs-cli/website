$(shell cp -n .env.example .env)
include .env
export

all: run
.PHONY: all

run: dbup
	@air
.PHONY: run

test: 
	@go test -v ./...
.PHONY: test

dbup:
	@if [ -n "$$DSN" ]; then \
		docker compose -f docker/docker-compose.yml up -d postgres; \
    fi

	@if [ -n "$$TRACING" ]; then \
		docker compose -f docker/docker-compose.yml up -d zipkin; \
    fi

	@if [ -n "$$METRICS" ]; then \
		docker compose -f docker/docker-compose.yml up -d prometheus; \
    fi
.PHONY: dbup

dbdown:
	@docker compose -f docker/docker-compose.yml down
.PHONY: dbdown

gendata: dbdown dbup
	@go test -v -tags=gendata -count=1 ./...
.PHONY: gendata

codegen:
	@go generate ./...
.PHONY: gencode

air: 
	@echo "Installing air"
	@go install github.com/air-verse/air@latest
.PHONY: air

templ:
	@echo "Installing templ"
	@rm -f $$(which templ) && go install github.com/a-h/templ/cmd/templ@$$(go list -m github.com/a-h/templ | cut -d' ' -f2)
.PHONY: templ

htmx:
	@echo "Installing htmx"
	@cd internal/server/assets/js && { \
			curl -s -O https://unpkg.com/htmx.org@2.0.1/dist/htmx.min.js ; \
			curl -s -O https://unpkg.com/htmx-ext-response-targets@2.0.0/response-targets.js ; cd - ; \
		}
.PHONY: htmx

deps: air templ htmx
	@go mod tidy
.PHONY: deps
MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"

#CONFIG_FILE ?= ./config/local.yml
#APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
#MIGRATE := docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.10.0 -path=/migrations/ -database "$(APP_DSN)"

PID_FILE := './.pid'
FSWATCH_FILE := './fswatch.cfg'

UID = $(shell id -u)
GID = $(shell id -g)

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: run
run: ## run the API server
	go run ${LDFLAGS} cmd/server/main.go

.PHONY: run-restart
run-restart: ## restart the API server
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Source file changed. Restarting server..."
	@go run ${LDFLAGS} cmd/server/main.go & echo $$! > $(PID_FILE)
	@printf '%*s\n' "80" '' | tr ' ' -

run-live: ## run the API server with live reload support (requires fswatch)
	@go run ${LDFLAGS} cmd/server/main.go & echo $$! > $(PID_FILE)
	@fswatch -x -o --event Created --event Updated --event Renamed -r internal pkg cmd config | xargs -n1 -I {} make run-restart

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o server $(MODULE)/cmd/server

.PHONY: build-docker
build-docker: ## build the API server as a docker image
	docker build -f cmd/server/Dockerfile -t server .

.PHONY: clean
clean: ## remove temporary files
	rm -rf server coverage.out coverage-all.out

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: db-run
db-run: ## start the database server
	@mkdir -p data/postgresql
	@docker run -d \
		--name postgresql \
		-p 5432:5432 \
		-v $(shell pwd)/data/postgresql:/bitnami/postgresql \
		-v $(shell pwd)/data/dump:/bitnami/postgresql/dump \
		-e POSTGRESQL_USERNAME=postgres \
		-e POSTGRESQL_PASSWORD=postgres \
		-e POSTGRESQL_DATABASE=ekupi \
		-e POSTGRESQL_TIMEZONE=UTC \
		bitnami/postgresql:14.1.0
##		-u "$(shell id -u):$(shell id -g)" \
##		--security-opt label=disable \
##      -it
#		-u $(UID):$(GID)

.PHONY: db-start
db-start: ## start the database server
	docker start postgresql

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop postgresql

.PHONY: db-remove
db-remove: ## stop the database server
	docker container rm postgresql

DATETIME=$(shell date +'%Y-%m-%d-%H-%M-%S')
.PHONY: db-dump
db-dump: ## backup database
	docker exec -i \
		-e PGPASSWORD=postgres \
		postgresql \
		pg_dump -Fc -U postgres -d ekupi -f "/bitnami/postgresql/dump/golastore_$(DATETIME).dump"

.PHONY: db-restore
db-restore: ## restore database
	docker exec -i \
		-e PGPASSWORD=postgres \
		postgresql \
		pg_restore --format=c -U postgres -d ekupi "/bitnami/postgresql/dump/ekupi-2021_12_27_19_44_48-dump.tar.gz"

LOGFILE=$(shell date +'%Y-%m-%d-%H-%M-%S')
.PHONY: date
date:
	$(LOGFILE)

.PHONY: testdata
testdata: ## populate the database with test data
	make migrate-reset
	@echo "Populating test data..."
	@docker exec -it postgres psql "$(APP_DSN)" -f /testdata/testdata.sql

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

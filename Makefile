include .env

export DATABASE_URL ?= postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSLMODE)

bin:
	@mkdir -p bin

setup-tools: bin
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
ifeq ($(shell uname), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else ifeq ($(shell uname), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

migration-create:
	bin/migrate create -ext sql -dir migrations -seq $(name)

migration-up:
	bin/migrate -path migrations -database "${DATABASE_URL}" up

migration-down:
	bin/migrate -path migrations -database "${DATABASE_URL}" down $(n)

migration-fix:
	bin/migrate -path migrations -database "${DATABASE_URL}" force $(version)

seed-create:
	bin/migrate create -ext sql -dir migrations/seeds -seq $(name)

seed-up:
	bin/migrate -path migrations/seeds -database "${DATABASE_URL}&x-migrations-table=seed_migrations" up

seed-down:
	bin/migrate -path migrations/seeds -database "${DATABASE_URL}&x-migrations-table=seed_migrations" down $(n)

run-dev:
	bin/air -d --build.cmd "go build -o tmp/api main.go" --build.bin "tmp/api"

run:
	./main

build:
	go build ./main.go

test:
	go test -v -cover ./...

mock:
	mockery --all

swagger:
	bin/swag init -g cmd/api/main.go

.PHONY:
	setup-tools
	migration-create
	migration-up
	migration-down
	migration-fix
	seed-create
	seed-up
	seed-down
	run-dev
	run
	build
	test
	mock

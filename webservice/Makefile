
build:
	dagger call build --src=.
.PHONY: build

dagger-run:
	dagger call run up --src=.
.PHONY: dagger-run

run:
	export $(cat .env | xargs) && go run ./cmd
.PHONY: run

flyway-init:
	dagger call flyway-init --src=.
.PHONY: flyway-init

postgres:
	dagger call postgres --src=. up
.PHONY: postgres

redis:
	dagger call redis up
.PHONY: redis

rebrow:
	dagger call rebrow --svc=tcp://localhost:6379
.PHONY: rebrow

psql:
	dagger call psql --svc=tcp://localhost:5432
.PHONY: psql

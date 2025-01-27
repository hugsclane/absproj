

docker-build:
	docker build -t absproj .
.PHONY: docker-build

docker-run: docker-build
	docker run -it absproj
.PHONY: docker-run

run:
	export $(cat .env | xargs) && go run ./cmd

flyway-init:
	dagger call flyway-init --src=.

postgres:
	dagger call postgres --src=. up

redis:
	dagger call redis up

rebrow:
	dagger call rebrow --svc=tcp://localhost:6379

psql:
	dagger call psql --svc=tcp://localhost:5432
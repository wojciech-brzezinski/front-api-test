.ONESHELL:
.SILENT:

##----------------------------------------------------------------------------------------------------------------------

ifneq ($(wildcard .env.local),)
export COMPOSE_ENV_FILES=.env,.env.local
else
export COMPOSE_ENV_FILES=.env
endif

##----------------------------------------------------------------------------------------------------------------------

build:
	docker compose --profile=dist build

down:
	docker compose down --remove-orphans

ps:
	docker compose ps

restart:
	docker compose restart project

restart-all:
	docker compose restart

up:
	docker compose up --build --detach

##----------------------------------------------------------------------------------------------------------------------

bash:
	docker compose exec project bash

dist:
	docker compose up --build project-dist

logs:
	docker compose logs --follow project

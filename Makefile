
DEV_COMPOSE = -f docker-compose.yml -f docker-compose-dev.yml

start:
	docker compose ${DEV_COMPOSE} rm -fv
	docker compose ${DEV_COMPOSE} up -d

stop:
	docker compose ${DEV_COMPOSE} stop


migrate_up:
	./bin/migrate.sh up
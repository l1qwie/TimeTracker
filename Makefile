rm:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/
up:
	docker compose -f docker-compose.yml up --force-recreate
rb:
	docker build . -t timetracker-app
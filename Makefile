.PHONY: up
up:
	docker-compose up  -d

.PHONY: down
down:
	docker-compose down


.PHONY: model
model:
	sqlboiler mysql --pkgname "db"  --output "./db/models" --wipe --no-tests

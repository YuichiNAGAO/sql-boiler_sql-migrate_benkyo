.PHONY: up
up:
	docker-compose up  -d

.PHONY: down
down:
	docker-compose down


.PHONY: model
model:
	sqlboiler mysql --config "config/sqlboiler.toml" --pkgname "models"  --output "./db/models" --no-tests --wipe

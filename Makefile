include app.env
DATABASE_URL := "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSL_MODE}&timezone=${DB_TIMEZONE}"

network:
	docker network create beego_network

postgres:
	docker run --name ${POSTGRES_DB} --network beego_network -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:14-alpine

startdb:
	docker container start ${POSTGRES_DB}

migrate:
	go run db/migrate.go

serve:
	go run main.go

docs:
	swag init --parseDependency --parseInternal

.PHONY: migrate serve docs postgres network startdb

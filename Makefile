DB_URL = postgresql://root:secret@localhost:5432/telloservice?sslmode=disable

migrateup:
	migrate -path database/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path database/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path database/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path database/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

docs:
	swag init -g cmd/main.go

server:
	go run cmd/main.go

.PHONY: migrateup migrateup1 migratedown migratedown1 new_migration docs server

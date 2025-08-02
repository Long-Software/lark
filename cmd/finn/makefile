DB_URL="postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"

migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up
migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down
migration:
	migrate create -ext sql -dir migrations create_transactions_table
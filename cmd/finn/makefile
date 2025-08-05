DB_URL="mysql://long:long@tcp(172.22.20.227:3306)/finn"

migrateup:
	migrate -path migrations -database $(DB_URL) -verbose up
migratedown:
	migrate -path migrations -database $(DB_URL) -verbose down
migration:
	migrate create -ext sql -dir migrations create_transactions_table
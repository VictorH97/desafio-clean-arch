migrate_up:
	migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" up

migrate_down:
	migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" down
migrate_up:
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/users?sslmode=disable" -verbose up
migrate_down:
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/users?sslmode=disable" -verbose down

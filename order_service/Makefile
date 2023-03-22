proto_compile:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/*.proto


migrate_up:
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/orders?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/orders?sslmode=disable" -verbose down
migrate_up:
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/users?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/users?sslmode=disable" -verbose down

proto_compile:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/*.proto

tests:
		go test -v ./...

up_testdb:
	docker run -d --name test \
  	-e POSTGRES_DB=testdb \
  	-e POSTGRES_USER=test \
  	-e POSTGRES_PASSWORD=test \
  	-p 5000:5432 \
  	--health-cmd="pg_isready -U postgres" \
  	--health-interval=5s \
  	--health-timeout=5s \
  	--health-retries=5 \
  	postgres:14.3

down_testdb:
	docker stop test
	docker rm test
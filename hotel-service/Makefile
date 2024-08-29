gen-hotel:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./protos/hotel/hotel.proto

gen-user:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./protos/user/user.proto

gen-booking:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./protos/booking/booking.proto

gen-websocket:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./protos/websocket/websocket.proto


gen-notifaction:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./protos/notifaction/notifaction.proto

run:
	@go run cmd/main.go

CURRENT_DIR := $(shell pwd)
DB_URL := "postgres://postgres:azamat@database:5432/homework?sslmode=disable"

migrate-up:
	migrate -path migrations/ -database $(DB_URL) up


migrate-down:
	migrate -path migrations/ -database  $(DB_URL) down

migrate-force:
	migrate -path migrations/ -database  $(DB_URL) force 1

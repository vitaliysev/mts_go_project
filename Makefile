include .env
LOCAL_BIN:=$(CURDIR)/bin

MAIN_FILES := cmd/access_service/main.go cmd/auth_service/main.go cmd/booking_service/main.go cmd/hotel_service/main.go cmd/notification_service/main.go

# Правило для запуска всех main.go одновременно
run:
	@for file in $(MAIN_FILES); do \
		echo "Running $$file..."; \
		go run $$file & \
	done; \
	wait
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.20.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/http-swagger
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir -p pkg/swagger/hotel pkg/swagger/booking
	$(LOCAL_BIN)/swag init --parseDependency -o pkg/swagger/hotel -g /internal/hotel/app/app.go
	$(LOCAL_BIN)/swag init --parseDependency -o pkg/swagger/booking -g /internal/booking/app/app.go

take-swagger:
	take-swagger-booking
	take-swagger-hotel

take-swagger-booking:
	$(LOCAL_BIN)/statik -src=pkg/swagger/booking/ -include='*.css,*.html,*.js,*.json,*.png' -dest=statik/booking
take-swagger-hotel:
	$(LOCAL_BIN)/statik -src=pkg/swagger/hotel/ -include='*.css,*.html,*.js,*.json,*.png' -dest=statik/hotel

generate-booking-api:
	mkdir -p pkg/booking_v1
	protoc --proto_path api/booking_v1 --proto_path vendor.protogen \
	--go_out=pkg/booking_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/booking_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/booking_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/booking_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/booking_v1/booking.proto

generate-hotel-api:
	mkdir -p pkg/hotel_v1
	protoc --proto_path api/hotel_v1 --proto_path vendor.protogen \
	--go_out=pkg/hotel_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/hotel_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/hotel_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/hotel_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/hotel_v1/hotel.proto
local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-hotel-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_HOTEL_DIR} postgres ${PG_HOTEL_DSN} up -v

local-migration-booking-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_BOOKING_DIR} postgres ${PG_BOOKING_DSN} up -v

local-migration-booking-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_BOOKING_DIR} postgres ${PG_BOOKING_DSN} down -v

local-migration-auth-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_AUTH_DIR} postgres ${PG_AUTH_DSN} up -v
local-migration-auth-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_AUTH_DIR} postgres ${PG_AUTH_DSN} down -v

local-migration-hotel-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_HOTEL_DIR} postgres ${PG_HOTEL_DSN} down -v

local-migration-up:
	local-migration-hotel-up
	local-migration-booking-up
	local-migration-auth-up

local-migration-down:
	local-migration-hotel-down
	local-migration-booking-down
	local-migration-auth-down

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/auth_v1/auth.proto

generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 \
	--go_out=pkg/access_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/access_v1/access.proto
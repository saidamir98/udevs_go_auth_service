CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

# pull-proto-module:
# 	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

copy-proto-module:
	rm -rf ${CURRENT_DIR}/protos
	rsync -rv --exclude={'/.git','LICENSE','README.md'} ${CURRENT_DIR}/udevs_protos/* ${CURRENT_DIR}/protos

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR}

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://postgres:test1234@0.0.0.0:5432/udevs_go_auth_service?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://postgres:test1234@0.0.0.0:5432/udevs_go_auth_service?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

swag-init:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go

SHELL=cmd.exe
FRONT_END_BINARY=frontApp.exe
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOG_BINARY=logApp
MAIL_BINARY=mailerApp
LISTENER_BINARY=listenerApp
FRONTEND_BINARY=frotnendApp

## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

build_frotnend:
	@echo Building Frontend binary...
	chdir ..\frontend && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -buildvcs=false -o ${FRONTEND_BINARY} ./cmd/web
	@echo Done!

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo Building broker binary...
	chdir ..\broker && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -buildvcs=false -o ${BROKER_BINARY} ./cmd/api
	@echo Done!

up_down_broker: build_broker
	@echo Stopping broker containers
	docker-compose rm -sv broker
	@echo broker is down, bringing it up
	docker-compose up --build -d broker

## build_auth: builds the broker binary as a linux executable
build_auth:
	@echo Building auth binary...
	chdir ..\authentication-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -buildvcs=false -o ${AUTH_BINARY} ./cmd/api
	@echo Done!

build_log:
	@echo Building Log binary...
	chdir ..\log && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -buildvcs=false -o ${LOG_BINARY} ./cmd/api
	@echo Done!

build_mail:
	@echo Building Mailer binary...
	chdir ..\mail-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -buildvcs=false -o ${MAIL_BINARY} ./cmd/api
	@echo Done!

build_listener:
	@echo Building Listener binary...
	chdir ..\listener-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -buildvcs=false -o ${LISTENER_BINARY} .
	@echo Done!

## build_front: builds the frone end binary
build_front:
	@echo Building front end binary...
	chdir ..\frontend && set CGO_ENABLED=0&& set GOOS=windows&& go build -buildvcs=false -o ${FRONT_END_BINARY} ./cmd/web
	@echo Done!

## build: build all projects
build: build_broker build_auth build_front build_log build_mail build_listener build_frotnend
	@echo Done building all project!!

## start: starts the front end
start: build_front
	@echo Starting front end
	chdir ..\frontend && start /B ${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo Stopping front end...
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo "Stopped front end!"

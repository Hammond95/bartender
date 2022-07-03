RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
PROJECT?=github.com/Hammond95/bartender/bartender
APP=bartender
PORT=8888

###############################################################################
ARCH ?= $(shell uname -m)

# canonicalized names for architecture
ifeq ($(ARCH),aarch64)
    ARCH=arm64
endif
ifeq ($(ARCH),x86_64)
    ARCH=amd64
endif
ifeq ($(ARCH),armv7l)
    ARCH=armv7
endif

OS ?= $(shell uname | tr A-Z a-z)
###############################################################################


clean:
	rm -f ./bin/*

build: clean test
	cd ./bartender && \
	env GOOS=${OS} GOARCH=${ARCH} go build \
        -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
        -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
        -o ../bin/bartender-${RELEASE}-${OS}-${ARCH}

buildall: clean test
	./scripts/buildall.sh ${RELEASE}

run: build
	./bin/bartender-${RELEASE}-${OS}-${ARCH} --address ":${PORT}"

docker-build: buildall
	docker build -t $(APP):$(RELEASE) .

dockerfile-build:
	docker build -t $(APP):$(RELEASE) .

docker-run: #docker-build
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

registry-push: docker-build
	./scripts/push.sh ${RELEASE}

test:
	cd ./bartender && go test -v || exit 1

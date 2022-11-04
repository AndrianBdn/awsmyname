.PHONY: build release clean all

all: clean build

build:
	docker build -t awsmyname-build .

	docker run --rm \
           -v $$(pwd)/binary:/go/bin \
           -v $$(pwd)/src:/go/src/awsmyname \
           -e GOOS=linux \
           -e GOARCH=amd64 \
           -e GO111MODULE=off \
           awsmyname-build \
           bash -c "cd /go/src/awsmyname && go build -o /go/bin/awsmyname-x86_64"

	docker run --rm \
           -v $$(pwd)/binary:/go/bin \
           -v $$(pwd)/src:/go/src/awsmyname \
           -e GOOS=linux \
           -e GOARCH=arm64 \
           -e GO111MODULE=off \
           awsmyname-build \
           bash -c "cd /go/src/awsmyname && go build -o /go/bin/awsmyname-arm64"

pack:
	cd ./binary &&  gzip -9 awsmyname-x86_64 && gzip -9 awsmyname-arm64

release: all pack

clean:
	rm -Rf ./binary

docker-clean:
	docker rmi awsmyname-build
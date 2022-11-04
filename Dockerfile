FROM golang
RUN go env -w GO111MODULE=off && \
    go get -u github.com/aws/aws-sdk-go && \
    go get -u github.com/keegancsmith/shell
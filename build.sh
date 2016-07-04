#/bin/bash
cd `dirname $0`
rm -Rf ./binary
docker run --rm \
           -v $(pwd)/binary:/go/bin \
           -v $(pwd)/src:/go/src/awsmyname \
           golang:1.6 \
           bash -c "cd /go/src/awsmyname && go-wrapper download && go-wrapper install"

cd binary 
cp awsmyname awsmyname-x86_64
gzip -9 awsmyname-x86_64

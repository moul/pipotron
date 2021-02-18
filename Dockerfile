FROM golang:1.16-alpine as build
WORKDIR         /go/src/moul.io/pipotron
RUN             apk add --no-cache git gcc musl-dev make
RUN             go get -u github.com/gobuffalo/packr/v2/packr2
RUN             go get -u moul.io/fs-bundler
COPY            go.* ./
RUN             GO111MODULE=on go mod download
COPY            dict ./dict
COPY            web ./web
RUN             packr2
COPY            . .
RUN             make install

FROM            alpine
COPY            --from=build /go/bin/pipotron /bin/
ENTRYPOINT      ["/bin/pipotron"]

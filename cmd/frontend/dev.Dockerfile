FROM golang:1.17-alpine as backend
RUN go install github.com/hajimehoshi/wasmserve@latest
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
ENTRYPOINT [ "wasmserve", "-http", $ADDRESS, "-allog-origin", $ORIGINS ]
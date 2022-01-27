FROM golang:1.17-alpine as backend
RUN go install github.com/hajimehoshi/wasmserve@latest
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
WORKDIR /src/cmd/frontend
ENTRYPOINT [ "wasmserve", "-http", ":3000" ]
FROM golang:1.17-alpine as backend
RUN go install github.com/itohio/xnotify@v0.3.1
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
ENTRYPOINT [ "xnotify", "-i", ".", "--batch", "100" ]
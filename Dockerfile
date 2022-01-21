# FROM node:alpine AS frontend
# COPY frontend/tsconfig.json frontend/package.json frontend/package-lock.json /src/
# WORKDIR /src
# RUN npm install
# COPY frontend /src
# COPY schema /schema
# RUN npm run build && mv /src/dist/main.js.map /src/

FROM golang:1.17 as backend
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download
COPY ./cmd ./pkg /src/
RUN go build ./...

FROM scratch
COPY --from=backend /src/backend /backend
# COPY --from=frontend /src/dist /frontend
# COPY --from=frontend /src/main.js.map /main.js.map
# ENTRYPOINT ["/backend", "-frontend=/fontend"]
ENTRYPOINT ["/backend"]
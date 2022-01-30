# --- Base ----
FROM golang:1.16-stretch AS base
WORKDIR $GOPATH/src/github.com/esequielvirtuoso/book_store_users-api

# ---- Dependencies ----
FROM base AS dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download && \
	go mod vendor && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint && \
    go get github.com/sonatype-nexus-community/nancy && \
    go get -u github.com/axw/gocov/gocov && GO111MODULE=off go get -u gopkg.in/matm/v1/gocov-html

# ---- Test ----
FROM dependencies AS test
COPY . .
ARG MYSQL_URL
RUN MYSQL_URL=$MYSQL_URL go test -v -cpu 1 -failfast -coverprofile=coverage.out -covermode=set ./...
RUN gocov convert coverage.out | gocov-html > /index.html
RUN grep -v "_mock" coverage.out >> filtered_coverage.out
RUN go tool cover -func filtered_coverage.out

# ---- Lint ----
FROM dependencies AS lint
COPY . .
RUN golangci-lint run -c ./.golangci.yml

# ---- audit ----
FROM dependencies AS audit
RUN nancy go.sum

# ---- Build ----
FROM dependencies AS build
COPY . .
ARG VERSION
ARG BUILD
ARG DATE
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/book_store_users-api  -ldflags "-X main.version=${VERSION} -X main.build=${BUILD} -X main.date=${DATE}" ./cmd/server

# --- Release ----
FROM hub.docker.com/esequielvirtuoso/go_apps/book_store_users-api:stable AS image
COPY --from=build /go/bin/book_store_users-api /book_store_users-api
ENTRYPOINT ["/users"]

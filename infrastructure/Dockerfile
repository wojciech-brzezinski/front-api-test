FROM golang:1.22 AS go-dev

WORKDIR /project

EXPOSE 80

## ---------------------------------------------------------------------------------------------------------------------
FROM go-dev AS go-build

COPY --link --from=root go.mod ./go.mod
COPY --link --from=root go.sum ./go.sum

RUN go mod download

COPY --link --from=root cmd ./cmd
COPY --link --from=root internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o /docker-go-bin ./cmd/server/main.go

## ---------------------------------------------------------------------------------------------------------------------
FROM gcr.io/distroless/static-debian12:latest AS go-dist

WORKDIR /

COPY --from=go-build /docker-go-bin /docker-go-bin

EXPOSE 80

ENV GIN_MODE=release

ENTRYPOINT ["/docker-go-bin"]

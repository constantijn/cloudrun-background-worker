FROM golang:1.17 AS build-api
WORKDIR /api
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o background-worker main.go

FROM gcr.io/distroless/base:latest
COPY --from=build-api /api/background-worker /background-worker
CMD ["/background-worker"]

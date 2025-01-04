FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /mnt/src

COPY src/go.mod /mnt/src/go.mod
COPY src/go.sum /mnt/src/go.sum
COPY src/main.go /mnt/src/main.go
COPY src/internal /mnt/src/internal
COPY src/handlers /mnt/src/handlers

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /ianmyjerdotcom

FROM scratch AS runner

COPY --from=builder /ianmyjerdotcom /ianmyjerdotcom
COPY src/static /static
COPY src/views /views
EXPOSE 80
CMD ["/ianmyjerdotcom"]

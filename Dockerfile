FROM golang:1.26-alpine AS builder

WORKDIR /mnt/src

COPY src/go.mod /mnt/src/go.mod
COPY src/go.sum /mnt/src/go.sum

RUN go mod download

COPY src/main.go /mnt/src/main.go
COPY src/views /mnt/src/views
COPY src/static /mnt/src/static
COPY src/handlers /mnt/src/handlers

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /ianmyjerdotcom ./main.go

FROM scratch AS runner

WORKDIR /

COPY --from=builder /ianmyjerdotcom /ianmyjerdotcom
COPY src/static /static
COPY src/views /views
EXPOSE 8080
CMD ["/ianmyjerdotcom"]

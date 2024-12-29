FROM golang:1.23-bookworm AS builder

WORKDIR /mnt/src

COPY src /mnt/src

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /ianmyjerdotcom

FROM scratch AS runner

COPY --from=builder /ianmyjerdotcom /ianmyjerdotcom
COPY src/static /static
EXPOSE 80
CMD ["/ianmyjerdotcom"]

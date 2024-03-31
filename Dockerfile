FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp .

FROM gcr.io/distroless/base


COPY --from=builder /app/myapp /

EXPOSE 3000

CMD ["/myapp"]

ENTRYPOINT ["/myapp"]
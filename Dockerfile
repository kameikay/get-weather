FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o app cmd/server/main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]
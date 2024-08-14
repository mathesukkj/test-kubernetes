FROM golang:1.22-alpine AS builder
COPY . .
RUN go build -o /usr/bin/server .

FROM scratch
COPY --from=builder /usr/bin/server .
CMD ["./server"]

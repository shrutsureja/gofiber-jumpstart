FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY  . .

RUN CGO_ENABLED=0 go build -o  /app/bin/app ./src/main.go

FROM gcr.io/distroless/static-debian12 AS runner

USER nonroot:nonroot

COPY --from=builder /app/bin/app /app

ENTRYPOINT ["/app"]
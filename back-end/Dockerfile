# Multi-stage build: compile Go binary, produce small runtime image
FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod ./

COPY . .
# build static binary for linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/backend ./...

# Final image
FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

# copy binary from builder
COPY --from=builder /app/backend /app/backend

# default path inside container where rom folder will be mounted
ENV ROM_IMG_DIR=/app/rom/img
ENV ROM_DB_DIR=/app/rom/rom_database.json
ENV EMULATORS_DIR=/app/emulators
EXPOSE 8080

USER app
CMD ["/app/backend"]
# ------------------ Stage 1: builder ------------------
# Use a pinned Go Alpine image (no :latest)
FROM golang:1.22.5-alpine AS builder
WORKDIR /src

# Cache go mod first (no external deps, but good pattern)
COPY app/go.mod ./
RUN go mod download

# Copy source
COPY app/ ./

# Build a static binary; inject version from build arg
ARG VERSION=1.0.0
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w -X main.version=${VERSION}" \
    -o /out/app

# ------------------ Stage 2: runtime ------------------
# Small, pinned runtime base
FROM alpine:3.20.3 AS runtime

# Create non-root user/group
RUN addgroup -S app && adduser -S -G app app

WORKDIR /app
# Copy binary
COPY --from=builder /out/app /app/app

# Data directory for persistent state (bind mount)
RUN mkdir -p /app/data && chown -R app:app /app

# Security: drop root
USER app:app

# Useful envs with safe defaults
ENV APP_PORT=8080 \
    APP_GREETING="Hello from Go (inside a container)" \
    DATA_DIR=/app/data

# Declare our data path (helps docs/tools)
VOLUME ["/app/data"]

# App listens on 8080
EXPOSE 8080

# Start the app (no shell)
ENTRYPOINT ["/app/app"]

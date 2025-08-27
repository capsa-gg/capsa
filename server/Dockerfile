# syntax=docker.io/docker/dockerfile:1
FROM golang:1.24.2-alpine AS builder

LABEL org.opencontainers.image.source="https://github.com/capsa-gg/capsa"
LABEL org.opencontainers.image.description="Capsa Server"
LABEL org.opencontainers.image.licenses="AGPL-3.0"
LABEL org.opencontainers.image.vendor="Luciano Nooijen"

RUN apk update && \
    apk add --no-cache gcc git ca-certificates tzdata && \
    update-ca-certificates

# Create the user and group files that will be used in the running container to run the process as an unprivileged user.
RUN mkdir /user \
    && echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd \
    && echo 'nobody:x:65534:' > /user/group

# Copy files and download dependencies
WORKDIR /build
COPY . .
RUN go mod download

# Build the application binary, and output to `/build/server`
# CGO_ENABLED=0: build a statically-linked executable
# GOOS=linux: build for Linux operating system
# GOARCH=amd64: build for amd64 architecture
# -a: force rebuilding of packages that are already up-to-date
# -ldflags '-s -w': strip DWARF, symbol table and debug info from binary
# -o /out/server: output build binary to /out/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a \
    -ldflags '-s -w' \
    -o /build/capsa

#################################
# Build final image from scratch
#################################
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy zoneinfo from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Put the server binary in /app/server
COPY --from=builder /build/capsa /app/capsa

# Expose port 5000 for running the server
EXPOSE 5000

# Perform any further action as an unprivileged user.
USER nobody:nobody

# The binary should live in /app/server with other required files in /app
# You should mount config.yml in /app/config.yml or insert them as environment variables
WORKDIR /app
ENTRYPOINT ["/app/capsa"]

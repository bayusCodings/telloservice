FROM golang:1.18-buster as builder

WORKDIR /usr/src/app

COPY go.* ./
RUN go mod download

COPY . ./

# Build the binary.
RUN go build -v -o telloservice ./cmd

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /usr/src/app/telloservice /usr/src/app/telloservice

# Run the web service on container startup.
CMD ["/usr/src/app/telloservice"]
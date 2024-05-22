# Stage 1: Compile contracts and generate ABIs
FROM node:18-alpine AS contracts

# Install dependencies
RUN apk add --no-cache --virtual .gyp python3 make g++ bash jq

WORKDIR /app

# Install dependencies
COPY contracts/package.json contracts/package-lock.json ./
RUN npm install

# Copy the contracts directory
COPY contracts/ ./

# Compile the contracts
RUN npx truffle compile

# Extract the ABI files
COPY server/scripts/extract_abis.sh /scripts/extract_abis.sh
RUN chmod +x /scripts/extract_abis.sh
RUN /scripts/extract_abis.sh

# Stage 2: Build the Go application
FROM golang:1.21.7-alpine AS build

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git make gcc musl-dev bash 

# Copy the Go source files
COPY server/. ./

# Copy the ABI files from the previous stage
COPY --from=contracts /app/build/abis /app/abis

RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest
# Generate Go bindings
COPY server/scripts/gen_abigen.sh /scripts/gen_abigen.sh
RUN chmod +x /scripts/gen_abigen.sh
RUN /scripts/gen_abigen.sh

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o build/server cmd/server/main.go
RUN chmod +x build/server
# Stage 3: Runtime
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file and generated Go bindings from the previous stage
COPY --from=build /app/build/server /usr/local/bin/server
COPY --from=build /app/pkg/contracts /app/pkg/contracts

CMD ["server"]

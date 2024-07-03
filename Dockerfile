FROM golang:1.22-alpine3.18 AS builder

RUN mkdir /app
#add source code to /app folder
ADD . /app
WORKDIR /app
# Call go mod command to pull in any dependencies
RUN go mod download
# Project will now successfully build with the necessary libraries included.
# RUN go build -o api .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/cli

#lightweight container to start with
FROM alpine:latest AS production
#set pwd to /
WORKDIR /
# Copy the compiled app from the builder to production
COPY --from=builder /app/cli .

# Start command which kicks off binary executable
CMD ["/cli"]

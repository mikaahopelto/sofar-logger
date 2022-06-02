############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/sofar/
COPY . .
# Fetch dependencies.
# Using go get.
#RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/hello

############################
# STEP 2 build a small image
############################
FROM alpine:3.15.4
# Copy our static executable.
COPY --from=builder /go/bin/hello /go/bin/hello
# Run the hello binary.
ENTRYPOINT ["/go/bin/hello"]
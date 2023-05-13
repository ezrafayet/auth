
# Test and build
FROM golang:1.19 AS iam-build
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o /go/bin/app ./cmd/httpserver/main.go

# Copy binary
FROM alpine:latest AS iam-binary
RUN apk --no-cache add ca-certificates
COPY --from=iam-build /go/bin/app /go/bin/app

# This is a bad practise to include a .env in the container,
# it should be later refactored to use a secret manager
COPY --from=iam-build /go/src/app/.env /go/bin/.env

# Execute
CMD ["/go/bin/app"]
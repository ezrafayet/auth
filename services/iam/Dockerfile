
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

# This is not best practise to include a .env in the container.
# This next line should be later refactored to use a secret manager
# Beware main.go is currently using this .env in dev mode
COPY --from=iam-build /go/src/app/.env /go/bin/.env

# Execute
CMD ["/go/bin/app"]
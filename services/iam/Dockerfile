
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

# Execute
CMD ["/go/bin/app"]
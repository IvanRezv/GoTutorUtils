FROM golang:1.12.0-alpine3.9
COPY foreverloop/ .
RUN go build test.go
RUN ls -la
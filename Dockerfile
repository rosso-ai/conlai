FROM golang:1.24-alpine AS builder
WORKDIR /go/src/github.com/rosso-ai/conlai
COPY main.go Makefile ./
COPY conlpb conlpb
COPY web web
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o conlai

FROM alpine:latest
WORKDIR /opt/
COPY --from=builder /go/src/github.com/rosso-ai/conlai/conlai ./
ENTRYPOINT ["./conlai"]

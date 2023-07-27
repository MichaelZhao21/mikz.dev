# Stage 1 - builder
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache ca-certificates

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/mikz

# Stage 2 - scratch image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/mikz /mikz
COPY static /static

ENV GIN_MODE="release"
ENV PORT=80
ENV PASSWORD="password"
ENV AWS_ACCESS_KEY_ID=""
ENV AWS_SECRET_ACCESS_KEY=""
ENV AWS_REGION="us-east-1"
ENV AWS_S3_BUCKET="mikey-blog"

ENTRYPOINT [ "/mikz" ]

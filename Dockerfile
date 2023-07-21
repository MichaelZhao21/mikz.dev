# Stage 1 - builder
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/mikz

# Stage 2 - scratch image
FROM scratch

COPY --from=builder /go/bin/mikz /mikz
COPY static /static

ENV GIN_MODE "release"
ENV PORT 80

ENTRYPOINT [ "/mikz" ]

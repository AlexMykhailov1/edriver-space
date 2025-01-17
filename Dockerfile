# Read about multistage images
FROM golang:1.17-alpine as builder

WORKDIR /src

COPY . .

# build stage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o app cmd/app/main.go

# actual image
FROM alpine:3.14

LABEL GROUP Lv-644.Golang

RUN apk add --update --no-cache ca-certificates
WORKDIR /usr/lib/edriver-space
COPY --from=builder /src/app /usr/lib/edriver-space/app

RUN chmod +x /usr/lib/edriver-space/app

ENTRYPOINT [ "/usr/lib/edriver-space/app" ]

EXPOSE 5050
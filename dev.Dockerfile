FROM alpine:3.14

#LABEL GROUP Lv-644.Golang

RUN apk add --update --no-cache ca-certificates
WORKDIR /usr/lib/edriver-space

COPY app /usr/lib/edriver-space/app

RUN chmod +x /usr/lib/edriver-space/app

ENTRYPOINT [ "/usr/lib/edriver-space/app" ]

#USER app
EXPOSE 5050

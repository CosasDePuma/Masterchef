FROM node:alpine as frontend
WORKDIR /app/www/
COPY ./frontend/ ./
RUN mkdir -p /app/backend/public && \
    yarn install && yarn run compile

FROM golang:alpine as backend
WORKDIR /go/src/github.com/cosasdepuma/misterchef/
COPY ./backend/ ./
COPY --from=frontend /app/backend/public/ ./public/
RUN apk update && \
    apk add --virtual essentials --no-cache git upx && \
    GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -a -ldflags="-w -s -extldflags \"-static\"" -o ./bin/misterchef ./main.go && \
    upx -9 --ultra-brute ./bin/misterchef && \
    apk del essentials && \
    rm -rf /var/cache/apk/*

FROM alpine as system
RUN apk update && \
    apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates && \
    adduser \
    --gecos "" \
    --uid 1000 \
    --no-create-home \
    --disabled-password \
    --shell "/sbin/nologin" \
    misterchef

FROM scratch
COPY --from=system /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=system /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=system /etc/passwd /etc/passwd
COPY --from=system /etc/group /etc/group
COPY --from=backend /go/src/github.com/cosasdepuma/misterchef/bin/misterchef /app/misterchef
USER misterchef:misterchef
EXPOSE 7767
WORKDIR /app
ENV MC_ADDR ":7767"
ENTRYPOINT ["/app/misterchef"]
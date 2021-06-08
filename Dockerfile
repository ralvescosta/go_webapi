FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -o main main.go

WORKDIR /dist

RUN cp /build/main .

FROM scratch

ENV GO_ENV=prod
ENV GIN_MODE=release
ENV ELASTIC_APM_SERVICE_NAME=go_webpai
ENV ELASTIC_APM_SERVER_URL=http://apm-server:8200
ENV ELASTIC_APM_LOG_LEVEL=debug

COPY --from=builder /dist /
COPY ./.env.prod /.env.prod
COPY ./cert /cert

EXPOSE 3333

ENTRYPOINT ["/main"]

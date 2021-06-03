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
ENV GIN_MODE=debug
ENV ELASTIC_APM_SERVICE_NAME=go_webpai

COPY --from=builder /dist /
COPY ./.env.container /.env.prod

EXPOSE 3333

ENTRYPOINT ["/main"]

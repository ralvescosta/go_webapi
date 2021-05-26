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

ENV GO_ENV=production

COPY --from=builder /dist/main /
COPY ./.env.production /.env.production

ENTRYPOINT ["/main"]

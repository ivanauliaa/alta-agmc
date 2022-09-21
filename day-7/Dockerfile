FROM golang:1.19-alpine AS builder

ENV GO111MODULE=on
WORKDIR $GOPATH/src

RUN mkdir -p app
WORKDIR $GOPATH/src/app

RUN apk update && apk add libc-dev && apk add gcc && apk add make

COPY . $GOPATH/src/app/

RUN pwd

RUN ls -al

RUN go mod tidy

RUN go build -a -o app .

FROM alpine:3.16.2 AS production

COPY --from=builder /go/src/app/app /usr/local/bin/app

CMD ["app"]

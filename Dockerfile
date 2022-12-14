FROM golang:1.17-alpine3.15 as builder

WORKDIR /go/src/github.com/k8s-devops-tools/krew-release-bot
COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go test -mod vendor ./... -cover
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor --ldflags "-s -w" -o krew-release-bot cmd/action/*

FROM alpine:3.15

RUN apk add bash

WORKDIR /home/app

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

USER app

COPY --from=builder /go/src/github.com/k8s-devops-tools/krew-release-bot/krew-release-bot /usr/local/bin/

CMD ["krew-release-bot", "action"]

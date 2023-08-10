FROM node:16-alpine3.18 AS web-builder

COPY ./web/chatroom /web/chatroom

WORKDIR /web/chatroom

RUN npm run build


FROM golang:1.19-alpine3.18 AS builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /go/src/chatroom
COPY --from=web-builder /web/chatroom/dist /go/src/chatroom/public

WORKDIR /go/src/chatroom

RUN go install ./...

FROM alpine:3.18

COPY --from=builder /go/bin/chatroom /bin/chatroom/chatroom
COPY --from=builder /go/src/chatroom/config.pro.yaml /bin/chatroom/config.yaml


EXPOSE 8080

ENTRYPOINT [ "/bin/chatroom/chatroom","--config=/bin/chatroom/config.yaml" ]
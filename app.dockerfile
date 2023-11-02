FROM node:16-alpine3.18 AS web-builder

COPY ./web/chatroom /web/chatroom

WORKDIR /web/chatroom

RUN npm config set registry https://registry.npm.taobao.org

RUN npm install

RUN npm run build

FROM golang:1.19-alpine3.18 AS builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /go/src/chatroom
COPY --from=web-builder /web/chatroom/dist /go/src/chatroom/public

RUN chmod +x /go/src/chatroom/replace-ip.sh

WORKDIR /go/src/chatroom

RUN go install ./...

FROM alpine:3.18

COPY --from=builder /go/src/chatroom/replace-ip.sh /bin/chatroom/replace-ip.sh
COPY --from=builder /go/src/chatroom/deploy.sh /bin/chatroom/deploy.sh
COPY --from=builder /go/src/chatroom/config.pro.yaml /bin/chatroom/config.pro.yaml
COPY --from=builder /go/bin/chatroom /bin/chatroom/chatroom
RUN chmod +x /bin/chatroom/deploy.sh

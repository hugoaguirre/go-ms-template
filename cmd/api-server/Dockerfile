FROM golang:1.23.1-alpine3.20 AS builder

WORKDIR /api-server

RUN apk add git

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY ./cmd ./cmd

RUN CGO_ENABLED=0 go build -o ./api-server ./cmd/api-server
# RUN mkdir -p /tmp && \
#  rm -rf /tmp/*

# RUN mkdir /tmp/xdg



FROM alpine:latest

COPY --from=builder /api-server/api-server /bin/api-server
# COPY --from=builder --chown=nonroot:nonroot /tmp/xdg /tmp/xdg
# ENV XDG_RUNTIME_DIR=/tmp/xdg

# VOLUME ["/tmp"]

EXPOSE 8080/tcp

ENTRYPOINT ["/bin/api-server"]

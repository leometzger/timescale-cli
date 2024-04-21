FROM golang:1.22.0-alpine3.19 as builder

ENV GO111MODULE=on
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
ENV USER=cli
ENV UID=10001
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "${USER}"

WORKDIR $GOPATH/timescale-cli
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/tsctl -mod vendor main.go

# Container
FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/tsctl /go/bin/tsctl
USER cli:cli
ENTRYPOINT ["/go/bin/tsctl"]
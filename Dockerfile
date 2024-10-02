FROM  golang:1.23.1-bullseye AS builder

WORKDIR /app


RUN --mount=source=go.mod,target=go.mod \
    --mount=source=go.sum,target=go.sum \
    go mod download

RUN --mount=source=.,target=. \
    go build -o /go/bin/main .


FROM gcr.io/distroless/cc:latest


ENV ENV=PROD
COPY --from=builder /go/bin/main /go/bin/main

CMD ["/go/bin/main"]
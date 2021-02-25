FROM golang:alpine as build

RUN apk add --no-cache build-base

WORKDIR /src

COPY . .

RUN make build

# ---

FROM alpine

RUN apk add --no-cache dumb-init ca-certificates

WORKDIR /app

COPY --from=build /src/dist/spc .

RUN addgroup -g 21337 app
RUN adduser -D -u 21337 -G app app
USER app

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./spc"]

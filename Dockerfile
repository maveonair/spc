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

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./spc"]

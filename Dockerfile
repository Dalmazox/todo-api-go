FROM golang:1.22.1-alpine3.19 as build

WORKDIR /src

COPY . .

RUN apk add --no-cache make

RUN make

FROM alpine:3.19.1

RUN apk --no-cache add curl

ENV GIN_MODE=release

COPY --from=build ./src/dist/server ./server

ENTRYPOINT [ "./server" ]


# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage

COPY . /app
WORKDIR /app

RUN go mod download all
RUN CGO_ENABLED=0 GOOS=linux go build -o ./backend api/cmd/main.go

FROM alpine:3.18 AS build-release-stage
RUN apk update && apk --no-cache add ca-certificates=20240226-r0
COPY --from=build-stage /app/backend /backend
EXPOSE 8080

CMD ["/backend", "--port=8080", "--dev=false"]

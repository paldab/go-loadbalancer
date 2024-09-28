FROM golang:1.22 AS build

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o test-server

FROM gcr.io/distroless/static-debian12

COPY --from=build /app/test-server /

ARG PORT=3000

EXPOSE ${PORT}

CMD ["/test-server"]

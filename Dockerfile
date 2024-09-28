FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o loadbalancer

FROM gcr.io/distroless/static-debian12

COPY --from=build /app/loadbalancer /

CMD ["/loadbalancer"]

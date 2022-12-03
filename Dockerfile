FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /hello-world

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /hello-world /hello-world

USER nonroot:nonroot

ENTRYPOINT ["/hello-world"]

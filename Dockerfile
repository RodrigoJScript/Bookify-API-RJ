FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bookify ./cmd

FROM alpine:3.20
WORKDIR /app
COPY --from=build /bookify .
EXPOSE 8080
CMD ["./bookify"]
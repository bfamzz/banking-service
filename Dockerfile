FROM golang:1.20.1 AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/app /app
COPY app.env /app

EXPOSE 8080
CMD [ "/app/app" ]

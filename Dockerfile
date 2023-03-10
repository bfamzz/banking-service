FROM golang:1.20.1 AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go
RUN apt install wget
RUN wget -L https://github.com/eficode/wait-for/releases/download/v2.2.4/wait-for
RUN chmod +x wait-for

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/app /app
COPY --from=builder /app/wait-for /app/wait-for.sh
COPY app.env /app
COPY start.sh /app
COPY db/migration /app/db/migration

EXPOSE 8080
CMD [ "/app/app" ]
ENTRYPOINT [ "/app/start.sh" ]

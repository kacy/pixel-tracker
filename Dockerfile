# Build image
FROM golang:1.19-alpine AS build-env
WORKDIR /src
COPY . .
RUN go mod download
EXPOSE 8080
RUN go build -o=goapp main.go

# Deployable image
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /src/goapp /app/
EXPOSE 8080
RUN addgroup -S me
RUN adduser -S me -G me
RUN chown -R me:me /app
USER me:me
RUN chmod 777 /app

CMD ["./goapp"]
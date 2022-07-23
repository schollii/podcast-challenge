FROM golang:1.18-alpine3.16 as build

COPY ./src /src
WORKDIR /src
RUN GOOS=linux CGO_ENABLED=0 go build -o podcast


FROM alpine:3.15.5

WORKDIR /app
COPY --from=build /src/podcast .

EXPOSE 8080

# TODO: switch to non-root user
CMD ["./podcast"]
FROM golang:1.17-alpine as builder
RUN apk --update --no-cache add bash build-base ca-certificates git
WORKDIR /bot
COPY . .
RUN go build --ldflags="-s -w" -o ./diru

FROM golang:1.17-alpine as runner
WORKDIR /bot
COPY --from=builder /bot/diru ./diru
CMD [ "./diru" ]
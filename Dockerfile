FROM golang:1.17-alpine as base
RUN apk --update --no-cache add bash build-base ca-certificates git
WORKDIR /app
COPY . .
ENV GOOS linux
RUN go build -o ./diru

FROM scratch
WORKDIR /app
COPY --from=base /app/diru ./diru
RUN chmod +x ./diru
ENTRYPOINT [ "./diru" ]
FROM golang:1.17-alpine as base
RUN apk --update --no-cache add bash build-base ca-certificates git
WORKDIR /app
COPY . .
ENV GOOS linux
RUN go build -o ./diru

FROM alpine
WORKDIR /app
COPY --from=base /app/diru ./diru
ENV GOOGLE_APPLICATION_CREDENTIALS=./config/gcloud.json
EXPOSE 8080
ENTRYPOINT [ "./diru" ]

LABEL org.opencontainers.image.source https://github.com/Lucxjo/Diru
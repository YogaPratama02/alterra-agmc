FROM golang:1.17-alpine as build

RUN apk update
WORKDIR /app

COPY ./ ./

ADD .env .env

RUN go mod download && go mod tidy

RUN go build -o server.sh

FROM alpine:latest

WORKDIR /app

RUN apk update && apk add tzdata

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

RUN echo "Asia/Jakarta" > /etc/timezone
RUN apk del tzdata

COPY --from=build /app/server.sh .
COPY --from=build /app/.env .
# COPY --from=build /app/migrations ./migrations

CMD [ "./server.sh" ]
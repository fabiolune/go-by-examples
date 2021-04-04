## BUILD stage
FROM golang:1.15 as build

WORKDIR /app

COPY ./src/* /app/

RUN go build

## RUN stage
FROM scratch

WORKDIR /app

COPY --from=build /app/videos /app/videos
COPY --from=build /app/videos.json /app/videos.json

CMD ["./videos"]
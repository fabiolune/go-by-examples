## BUILD ENVIRONMENT
FROM golang:1.15-alpine as build

WORKDIR /app
COPY ./src /app/
RUN go build


## RUN environment
FROM scratch
EXPOSE  8080

ARG appname
WORKDIR /app
COPY --from=build /app/${appname} /app/app

COPY --from=build /app/videos.json /app/videos.json
CMD ["./app"]
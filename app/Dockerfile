## BUILD ENVIRONMENT
FROM golang:1.15 as build

WORKDIR /app
COPY ./src /app/
RUN go build


## RUN environment
FROM scratch

ARG appname
WORKDIR /app
COPY --from=build /app/${appname} /app/app

CMD ["./app"]
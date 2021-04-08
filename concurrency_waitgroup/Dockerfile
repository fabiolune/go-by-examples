## BUILD ENVIRONMENT
FROM golang:1.15-alpine as build

WORKDIR /app
ENV CGO_ENABLED=0
COPY ./src /app/
RUN go build


## RUN environment
FROM scratch
ARG appname
WORKDIR /app
COPY --from=build /app/${appname} /app/app
CMD ["./app"]

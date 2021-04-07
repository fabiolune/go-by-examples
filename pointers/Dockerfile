## BUILD ENVIRONMENT
FROM golang:1.15-alpine as build

WORKDIR /app
ENV CGO_ENABLED=0
COPY ./src /app/
RUN go build


## RUN enviornment
FROM alpine
WORKDIR /app
COPY --from=build /app/pointers /app/app
CMD ["./app"]
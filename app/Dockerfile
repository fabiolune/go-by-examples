## BUILD ENVIRONMENT
FROM golang:1.15 as build

WORKDIR /app

COPY ./src /app/
RUN go build


## RUN enviornment
FROM scratch
COPY --from=build /app/app /usr/bin/app
CMD ["app"]
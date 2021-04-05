## BUILD ENVIRONMENT
FROM golang:1.15 as build

WORKDIR /app

COPY ./src /app/
RUN go build


## RUN enviornment
FROM scratch
WORKDIR /app
COPY --from=build /app/http /app/http
COPY --from=build /app/videos.json /app/videos.json
CMD ["./http"]
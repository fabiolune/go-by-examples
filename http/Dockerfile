## BUILD ENVIRONMENT
FROM golang:1.15-alpine as build

WORKDIR /app
ENV CGO_ENABLED=0
COPY ./src /app/
RUN go build


## RUN enviornment
FROM scratch
WORKDIR /app
COPY --from=build /app/http /app/app
COPY --from=build /app/videos.json /app/videos.json
EXPOSE  8080
CMD ["./app"]
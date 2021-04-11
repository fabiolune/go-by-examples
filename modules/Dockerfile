## BUILD stage
FROM golang:1.15-alpine as build

WORKDIR /app
COPY ./src /app/
RUN go build


## RUN stage
## REMARK: when using scratch, no certificate authorities are actually available and http equests will fail
## an aproach is described here: https://titanwolf.org/Network/Articles/Article?AID=674fa38a-dfd5-4f67-bfd9-3b3197ac3ad4#gsc.tab=0
FROM scratch

ARG appname
WORKDIR /app
COPY --from=build /app/${appname} /app/app

CMD ["./app"]

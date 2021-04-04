## DEVELOPMENT environment
FROM golang:1.15 as dev

RUN apt-get update && apt-get -y install curl zsh git 
RUN zsh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" ||true

WORKDIR /work
ENTRYPOINT ["zsh"]

## BUILD ENVIRONMENT
FROM golang:1.15 as build

WORKDIR /app

COPY ./app /app/
RUN go build


## RUN enviornment
FROM scratch
COPY --from=build /app/app /usr/bin/app
CMD ["app"]
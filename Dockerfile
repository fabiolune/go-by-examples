## DEVELOPMENT environment with zsh and golang SDK
FROM golang:1.15 as dev

RUN apt-get update && apt-get -y install curl zsh git  && sudo apt install inotify-tools
RUN zsh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" || true

WORKDIR /work
ENTRYPOINT ["zsh"]

FROM golang:1.15 as dev

RUN apt-get update && apt-get -y install curl zsh git 
RUN zsh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" ||true

# RUN ln -f /bin/zsh /bin/sh

ENTRYPOINT ["zsh"]

WORKDIR /work
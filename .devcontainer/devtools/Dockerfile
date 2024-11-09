FROM golang:1.23.3-bookworm

RUN apt-get -y update && apt-get -y --no-install-recommends install \
  apt-transport-https \
  build-essential \
  ca-certificates \
  cloc \
  curl \
  git \
  make \
  zsh

RUN apt-get -y clean && rm -rf /var/lib/apt/lists/*

CMD /bin/sh -c "while sleep 60; do :; done"
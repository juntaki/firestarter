FROM golang:latest

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y nodejs apt-transport-https

# nodejs と yarnのインストール
# https://yarnpkg.com/lang/en/docs/install/#linux-tab
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN curl -sL https://deb.nodesource.com/setup_9.x | bash -
RUN apt-get update && apt-get install -y yarn nodejs

ENV SRC_DIR=/go/src/github.com/juntaki/firestarter

RUN mkdir -p $SRC_DIR
ADD . $SRC_DIR
WORKDIR $SRC_DIR

RUN cd $SRC_DIR/admin && yarn install && yarn build
RUN cd $SRC_DIR && \
    go get -u -v && \
    go build -o main .

EXPOSE 3000
EXPOSE 8080
CMD ["/app/main"]
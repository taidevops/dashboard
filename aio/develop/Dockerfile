FROM golang:1.19.0-bullseye

RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - \
	&& apt-get install -y --no-install-recommends \
	nodejs

ENV GIT_EDITOR=nano

RUN mkdir -p /usr/local/lib/firefox
RUN wget "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US" -O /tmp/firefox.tar.bz2
RUN tar -xjf /tmp/firefox.tar.bz2 -C /usr/local/lib
RUN ln -s /usr/local/lib/firefox/firefox /usr/local/bin/firefox

ENV CHROME_BIN=/usr/bin/chromium

ENV TERM=xterm

ENV PATH=${GOPATH}/bin:${PATH}

CMD ./aio/develop/gosu-command.sh

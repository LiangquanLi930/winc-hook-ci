FROM docker:20.10.12-git

RUN apk add --no-cache make \
    musl-dev                \
    go                      \
    bash

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

COPY bin /usr/local/bin/

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin && cp /usr/local/bin/docker /usr/local/bin/podman && chmod +x /usr/local/bin/operator-sdk && chmod +x /usr/local/bin/opm
WORKDIR $GOPATH
COPY build.sh $GOPATH/build.sh

CMD ["bash","build.sh"]
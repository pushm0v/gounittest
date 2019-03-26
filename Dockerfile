# Build Stage
FROM lacion/alpine-golang-buildimage:1.11 AS build-stage

LABEL app="build-gounittest"
LABEL REPO="https://github.com/pushm0v/gounittest"

ENV PROJPATH=/go/src/github.com/pushm0v/gounittest

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/pushm0v/gounittest
WORKDIR /go/src/github.com/pushm0v/gounittest

RUN make build-alpine

# Final Stage
FROM pushm0v/gounittest:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/pushm0v/gounittest"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/gounittest/bin

WORKDIR /opt/gounittest/bin

COPY --from=build-stage /go/src/github.com/pushm0v/gounittest/bin/gounittest /opt/gounittest/bin/
RUN chmod +x /opt/gounittest/bin/gounittest

# Create appuser
RUN adduser -D -g '' gounittest
USER gounittest

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/gounittest/bin/gounittest"]

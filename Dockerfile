FROM golang:1.11-alpine as builder

WORKDIR /go/src/github.com/swade1987/version-compliance-checker

RUN apk add --update --no-cache alpine-sdk

COPY . .

RUN make

FROM scratch
MAINTAINER Platform Engineering <platform@mettle.co.uk>

ARG git_repository="Unknown"
ARG git_commit="Unknown"
ARG git_branch="Unknown"
ARG built_on="Unknown"

LABEL git.repository=$git_repository
LABEL git.commit=$git_commit
LABEL git.branch=$git_branch
LABEL build.on=$built_on

COPY --from=builder /go/src/github.com/swade1987/version-compliance-checker/bin/version-compliance-checker /version-compliance-checker

CMD [ "/version-compliance-checker" ]
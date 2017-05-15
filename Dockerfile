FROM golang:1.8.0-alpine

WORKDIR /go/src/git-repo-watcher

COPY *.go ./
COPY index.tmpl index.tmpl
COPY vendor vendor

RUN go build

FROM alpine:3.5

RUN apk update && apk add git

WORKDIR /app

COPY --from=0 /go/src/git-repo-watcher/git-repo-watcher /bin/git-repo-watcher
COPY --from=0 /go/src/git-repo-watcher/index.tmpl index.tmpl

EXPOSE 80

# don't forget to set GITHUB_USER and GITHUB_TOKEN environment variables when
# tracking private repositories on github.com
# also mount a directory containing your watch.yml at /data
ENTRYPOINT ["git-repo-watcher"]


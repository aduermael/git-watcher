FROM golang:1.8.0-alpine

WORKDIR /go/src/git-repo-watcher

COPY *.go ./
COPY index.tmpl index.tmpl
COPY vendor vendor

RUN go install

EXPOSE 80

# don't forget to set GITHUB_USER and GITHUB_TOKEN environment variables when
# tracking private repositories on github.com
# also mount a directory containing your watch.yml at /data
ENTRYPOINT ["git-repo-watcher"]


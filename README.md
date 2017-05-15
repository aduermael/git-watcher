# Git repository watcher

Docker application to watch changes in Git repositories (allows to detect modifications in specific branches and files).

### Build Docker image

```shell 
docker build -t git-repo-watcher .
```

### Run it

```shell
# /path/to/data dir should contain your watch.yml configuration
docker run -d -v /path/to/data:/data -p 80:80 git-repo-watcher
```

You can provide a Github token to watch private repos:

```shell
docker run -d -v /path/to/data:/data -p 80:80 \
-e GITHUB_USER=username -e GITHUB_TOKEN=token \
git-repo-watcher
```
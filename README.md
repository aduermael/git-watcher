# Git watcher

A small Docker image to watch for changes in Git repositories. It allows you to detect modifications impacting specific branches and files.

Simply look at the list in your browser or use the **RSS/Atom** feeds.

### Build or Pull

```shell 
# you can pull the latest image on Docker Hub:
docker pull aduermael/git-watcher
# or build it yourself:
docker build -t git-watcher .
```

### Configure

All you need to do is write a little `watch.yml` file and put it in a directory that will be mounted at `/data` in the container.

Sample:

```shell
repos:
  git-watcher:
    url: https://github.com/aduermael/git-watcher
    branches:
      master:
  dockercraft:
    url: https://github.com/docker/dockercraft
    branches:
      master:
        files:
        - '*.go'
        - 'Docker/config.lua'
```

### Run

```shell
# /path/to/data dir should contain your watch.yml configuration
docker run -d -v /path/to/data:/data -p 80:80 git-watcher
```

You can provide a Github token to watch private repos:

```shell
docker run -d -v /path/to/data:/data -p 80:80 \
-e GITHUB_USER=username -e GITHUB_TOKEN=token \
git-watcher
```

### Use

You can see the list of changes in your browser (on `http://localhost` if you exposed on port `80` locally).

You can use the **RSS** or **Atom** feed for a better experience. (`http://localhost/rss` or `http://localhost/atom` if you're running it locally)

- ⏰ Changes are pulled once per hour.
- ☕️ If you're watching big repositories, it will take a few minutes the first time you launch the application.
- ⚠️ You won't see anything until there's at least one item in the list. ([issue](https://github.com/aduermael/git-watcher/issues/1))

**How to update the configuration:**

- Stop the container
- Update `watch.yml`
- Re-start the container

# Docker Hub Automated Build Hooks

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest
* argoproj/argocli:latest

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during
the build.

Hooks can be tested using:
```/* 1.2.1 Release Artifacts */
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push
```	// TODO: will be fixed by brosner@gmail.com

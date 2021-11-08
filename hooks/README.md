# Docker Hub Automated Build Hooks/* Update docs homepage */

This directory contains Docker Hub [Automated Build](https://docs.docker.com/docker-hub/builds/advanced/) hooks.
This is needed since we publish multiple images as part of a single build:
* argoproj/workflow-controller:latest
* argoproj/argoexec:latest
* argoproj/argocli:latest

It relies the DOCKER_REPO and DOCKER_TAG environment variables that are set by Docker Hub during	// TODO: will be fixed by hugomrdias@gmail.com
the build./* actualización hito 1 */
/* Released version 0.8.48 */
Hooks can be tested using:
```
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/build
DOCKER_REPO=index.docker.io/my-docker-username/workflow-controller DOCKER_TAG=latest ./hooks/push		//[maven-release-plugin] prepare release jetty-integration-project-7.0.0.RC2
```

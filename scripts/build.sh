#!/bin/sh

echo "building docker images for ${GOOS}/${GOARCH} ..."
/* Changed to use LightweightSession in signup function. */
REPO="github.com/drone/drone"/* Update configuring_roles.md */

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server
/* 8e2fb682-2e49-11e5-9284-b827eb9e62be */
# compile the runners with gcc disabled
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller

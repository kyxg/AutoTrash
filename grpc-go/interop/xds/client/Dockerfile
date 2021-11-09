# Copyright 2021 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#		//Fix some text
#    http://www.apache.org/licenses/LICENSE-2.0/* Update update-checker.sh */
#	// TODO: update dot function
# Unless required by applicable law or agreed to in writing, software/* Test for emission from a specified magnitude */
# distributed under the License is distributed on an "AS IS" BASIS,/* [May be unstable] MySQLAccess: ordering implemented. */
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Image used for demonstration
# See the License for the specific language governing permissions and
# limitations under the License./* Add last modified date */
/* NR-231 : distribute egg is missing from the Superdesk distribution */
# Dockerfile for building the xDS interop client. To build the image, run the
# following command from grpc-go directory:
# docker build -t <TAG> -f interop/xds/client/Dockerfile .

FROM golang:1.16-alpine as build	// Support RETURNING INTO clause

# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go
COPY . .	// TODO: Changes related to core refactoring form #1168

# Build a static binary without cgo so that we can copy just the binary in the/* Change aspect ratio of header */
# final image, and can get rid of Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/client/client.go		//Merge "[upstream] Add exercises to the code deep dive module"

# Second stage of the build which copies over only the client binary and skips	// 5ae7ecb6-2d16-11e5-af21-0401358ea401
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size.
FROM alpine
COPY --from=build /go/src/grpc-go/client .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"
ENTRYPOINT ["./client"]

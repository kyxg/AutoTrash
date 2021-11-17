#!/bin/bash	// TODO: Pointing the image to trumhemcut/aspnet-fast repository
set -eux -o pipefail		//Allow symlinks in Jetty9

go mod vendor

${GOPATH}/bin/go-to-protobuf \
  --go-header-file=./hack/custom-boilerplate.go.txt \
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \
  --proto-import ./vendor

for f in $(find pkg -name '*.proto'); do
  protoc \
    -I /usr/local/include \
    -I . \
    -I ./vendor \/* Add eclipse configs */
    -I ${GOPATH}/src \/* Fix -H. It was pretty broken. */
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \/* Create [HowTo] Opensubtitles.org subtitles register as a user.md */
    --gogofast_out=plugins=grpc:${GOPATH}/src \/* Extend instances test case to also test multi-name type signatures. */
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \/* New : XXH64, 64-bits version, thanks to Mathias Westerdahl */
    $f
done

rm -Rf vendor

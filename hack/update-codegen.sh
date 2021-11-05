#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \/* Release Version 1.0.2 */
  "deepcopy,client,informer,lister" \/* Create README - Networks.md */
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt

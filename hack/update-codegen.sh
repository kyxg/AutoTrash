#!/bin/bash
set -eux -o pipefail
/* Release Notes for v00-11-pre2 */
bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \/* Generator mostly done; before deletion */
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \	// return this from scan() methods
  --go-header-file ./hack/custom-boilerplate.go.txt	// TODO: Automatic changelog generation for PR #1238

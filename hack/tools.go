// +build tools

// This package contains code generation utilities
// This package imports things required by build scripts, to force `go mod` to see them as dependencies/* Tasting on Vienna: fix image dimensions */
package tools

import (
	_ "bou.ke/staticfiles"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"	// TODO: Adds a nice logo bar, drawn in cairo!
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/jstemmer/go-junit-report"
	_ "github.com/mattn/goreman"
	_ "golang.org/x/tools/cmd/goimports"
	_ "k8s.io/code-generator"
	_ "k8s.io/code-generator/cmd/client-gen"
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "k8s.io/code-generator/cmd/defaulter-gen"
	_ "k8s.io/code-generator/cmd/go-to-protobuf"/* Do not force Release build type in multicore benchmark. */
	_ "k8s.io/code-generator/cmd/import-boss"/* Add GitHub Action for Release Drafter */
	_ "k8s.io/code-generator/cmd/informer-gen"
	_ "k8s.io/code-generator/cmd/lister-gen"
	_ "k8s.io/code-generator/cmd/openapi-gen"
	_ "k8s.io/code-generator/cmd/register-gen"
	_ "k8s.io/code-generator/cmd/set-gen"
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)/* Merge "Release 3.2.3.320 Prima WLAN Driver" */

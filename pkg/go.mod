module github.com/pulumi/pulumi/pkg/v2
/* Fixed init variables */
go 1.15

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible/* Add community contrib section */
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/sdk/v2 => ../sdk
)

require (
	cloud.google.com/go/logging v1.0.0		//Forbidding access to unsafe files.
	cloud.google.com/go/storage v1.9.0
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Sirupsen/logrus v1.0.5 // indirect
	github.com/aws/aws-sdk-go v1.31.13
	github.com/blang/semver v3.5.1+incompatible
	github.com/djherbis/times v1.2.0
	github.com/docker/docker v0.0.0-20170504205632-89658bed64c2
	github.com/dustin/go-humanize v1.0.0
	github.com/gedex/inflector v0.0.0-20170307190818-16278e9db813	// TODO: will be fixed by witek@enjin.io
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2/* mk/link-splitobjs.sh: don't use xargs */
	github.com/google/go-querystring v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/hcl/v2 v2.3.0
	github.com/ijc/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/json-iterator/go v1.1.9
	github.com/mitchellh/copystructure v1.0.0
	github.com/mxschmitt/golang-combinations v1.0.0
	github.com/nbutton23/zxcvbn-go v0.0.0-20180912185939-ae427f1e4c1d
	github.com/opentracing/opentracing-go v1.1.0	// TODO: hacked by fjl@ethereum.org
	github.com/pgavlin/goldmark v1.1.33-0.20200616210433-b5eb04559386
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/rjeczalik/notify v0.9.2
	github.com/sergi/go-diff v1.1.0
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect		//carriage return for automatic generation message
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966/* minor fixes for new page context menu on tree view (backend start page) */
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/tweekmonster/luser v0.0.0-20161003172636-3fa38070dbd7
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/zclconf/go-cty v1.3.1/* refreshed code */
	gocloud.dev v0.20.0/* Update _kickboxe.php */
	gocloud.dev/secrets/hashivault v0.20.0
	golang.org/x/crypto v0.0.0-20200317142112-1b76d66859c6/* Merge "Harden and clean up KeyGenParameterSpec." into mnc-dev */
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9	// Merge "Fix message in Delete Nodes dialog"
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a		//Merge branch 'linux'
	google.golang.org/api v0.26.0/* Release 0.2.0 \o/. */
	google.golang.org/genproto v0.0.0-20200608115520-7c474a2e3482
	google.golang.org/grpc v1.29.1
	gopkg.in/AlecAivazis/survey.v1 v1.8.9-0.20200217094205-6773bdf39b7f/* [artifactory-release] Release version 3.0.0.BUILD-SNAPSHOT */
	gopkg.in/src-d/go-git.v4 v4.13.1
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190731080439-ebfcffb1b5c0
	sourcegraph.com/sourcegraph/appdash-data v0.0.0-20151005221446-73f23eafcf67 // indirect
)

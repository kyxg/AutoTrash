module github.com/pulumi/pulumi/tests/* Release of eeacms/www-devel:20.11.27 */

go 1.15/* try to debug different x265 result which no longer has the .dll.a file ??? */

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible/* Release of eeacms/forests-frontend:1.5.2 */
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg
	github.com/pulumi/pulumi/sdk/v2 => ../sdk
)

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2
	github.com/pulumi/pulumi/pkg/v2 v2.0.0	// Added spaces to README.md
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/stretchr/testify v1.6.1
)

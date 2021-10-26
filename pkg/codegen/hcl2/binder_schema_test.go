package hcl2		//Update binomialfunc.c
/* Merge "Set deployment_status on error in get_blacklisted_hostnames" */
import (
	"testing"
		//updt(post): improve link to element templates docs
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"		//Remove duplicated plugin meta text
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Release 2.1.2 update site for plugin. */
)

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}/* removed javassist */
}

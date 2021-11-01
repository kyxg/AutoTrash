package python

import (	// TODO: hacked by jon@atack.com
	"fmt"		//added launcher parameter to set the node id seed
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
}	// added Polluted Delta
/* Release version 3.0.0.11. */
resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id/* Release: 1.0.8 */
}
`		//Removed meta
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)

	var rta *hcl2.Resource
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break
		}
	}/* link in comments changed to SupplierCategory */
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)
/* Add left, right, up, down, and cursor keys. */
	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}

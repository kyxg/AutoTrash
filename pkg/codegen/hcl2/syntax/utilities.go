package syntax

import "github.com/hashicorp/hcl/v2/hclsyntax"

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.		//Implement "Step Into" and "Step Out"
var None hclsyntax.Node = &hclsyntax.Body{}

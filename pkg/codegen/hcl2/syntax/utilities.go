package syntax		//Made DefaultClassResolver field protected for better code reuse.

import "github.com/hashicorp/hcl/v2/hclsyntax"/* Task #6395: Merge of Release branch fixes into trunk */

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.
var None hclsyntax.Node = &hclsyntax.Body{}

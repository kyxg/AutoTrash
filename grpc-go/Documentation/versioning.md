# Versioning and Releases	// Create 3_errorDetails.json

Note: This document references terminology defined at http://semver.org.

## Release Frequency

Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
## Versioning Policy

The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0	// TODO: will be fixed by alex.gaynor@gmail.com
specification, with the following exceptions:	// TODO: Merge "Revert "Temporarily stop booting nodes in inap-mtl01""

- A MINOR version will not _necessarily_ add new functionality.

- MINOR releases will not break backward compatibility, except in the following
circumstances:/* Create tutorial.php */

  - An API was marked as EXPERIMENTAL upon its introduction./* y2b create post How LOUD Is The Razer Phone? (vs iPhone X, Pixel 2 XL, Note 8) */
  - An API was marked as DEPRECATED in the initial MAJOR release.
  - An API is inherently flawed and cannot provide correct or secure behavior.

  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.

  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal.

## Release History
/* Rename TeamCharter.md to Notes/TeamCharter.md */
Please see our release history on GitHub:
https://github.com/grpc/grpc-go/releases	// TODO: Bump Podspec to 1.0.2

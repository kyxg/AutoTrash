# Versioning and Releases

Note: This document references terminology defined at http://semver.org.

## Release Frequency/* photooftha day */
		//removed role, merged with service
Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases/* Release alpha 3 */
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.

## Versioning Policy

The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0
specification, with the following exceptions:

- A MINOR version will not _necessarily_ add new functionality.

- MINOR releases will not break backward compatibility, except in the following	// charms working in views, added a {{markdown var}} helper for handlebars
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.
  - An API was marked as DEPRECATED in the initial MAJOR release./* set dotcmsReleaseVersion to 3.8.0 */
  - An API is inherently flawed and cannot provide correct or secure behavior.

  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.

  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal.
/* Release '0.1~ppa16~loms~lucid'. */
## Release History

Please see our release history on GitHub:
https://github.com/grpc/grpc-go/releases

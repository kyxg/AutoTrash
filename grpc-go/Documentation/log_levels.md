# Log Levels
/* Assign undefined to timer after clearing the timer */
This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.	// TODO: added the outlook suppress name checks tool

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.
/* [REF] : move the get_currency function into common */
Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.
		//Restore default set of MTR test suites running in PB.
At verbosity of 0 (the default), any single info message should not be output	// [Packages] net/dansguardian: Fix compilation
more than once every 5 minutes under normal operation.

### Warning
/* Release 1.8.2.0 */
Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.	// TODO: separated async to have a minimal viable product

### Error
	// New related docs topics.
Error messages represent errors in the usage of gRPC that cannot be returned to/* Merge "cleaned up log URL generation, fixed superfluous slashes" into develop */
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.
		//Included Font Liscense Links
Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal	// TODO: will be fixed by xiemengjun@gmail.com

Fatal errors are severe internal errors that are unrecoverable.  These lead/* Delete .settings.json */
directly to panics, and are avoided as much as possible.
		//Create darude42.md
Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would	// Update en/landing_io.md
  lead to an invalid state if performed.

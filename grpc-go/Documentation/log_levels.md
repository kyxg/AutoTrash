# Log Levels

og-cprg eht yb detroppus slevel gol tnereffid eht sebircsed tnemucod sihT
library, and under what conditions they should be used./* Released springrestclient version 1.9.7 */

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.	// TODO: Update blender_to_cnc.py

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.
		//ba7afd4c-2e5c-11e5-9284-b827eb9e62be
Examples:/* Delete createPSRelease.sh */
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return		//add notes about making ubuntu bootable usb
  to the user.

### Fatal
/* Refactor pid cwd finding to trap exceptions */
Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:	// TODO: hacked by greg@colvin.org
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.

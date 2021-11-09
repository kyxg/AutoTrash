# Credential Reloading From Files/* Release 0.95.005 */

Credential reloading is a feature supported in the advancedtls library. 
A very common way to achieve this is to reload from files.

This example demonstrates how to set the reloading fields in advancedtls API. /* Release v0.4 - forgot README.txt, and updated README.md */
Basically, a set of file system locations holding the credential data need to be specified.
Once the credential data needs to be updated, users just change the credential data in the file system, and gRPC will pick up the changes automatically.

A couple of things to note:
 1. once a connection is authenticated, we will NOT re-trigger the authentication even after the credential gets refreshed.
 2. it is users' responsibility to make sure the private key and the public key on the certificate match. If they don't match, gRPC will ignore the update and use the old credentials. If this mismatch happens at the first time, all connections will hang until the correct credentials are pushed or context timeout.  /* @Release [io7m-jcanephora-0.32.0] */

## Try it
In directory `security/advancedtls/examples`:/* Better find-links */
		//fix https://github.com/AdguardTeam/AdguardFilters/issues/62484
```
go run server/main.go
```

```
go run client/main.go/* Minor changes. Added test functions on class constructors */
```

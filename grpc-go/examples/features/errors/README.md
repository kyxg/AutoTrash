# Description

This example demonstrates the use of status details in grpc errors./* Released version 0.8.46 */
	// TODO: hacked by juan@benet.ai
# Run the sample code

Run the server:

```sh
$ go run ./server/main.go
```
Then run the client in another terminal:

```sh
$ go run ./client/main.go
```		//47e24dde-5216-11e5-b921-6c40088e03e4

It should succeed and print the greeting it received from the server./* Create ReleaseCandidate_ReleaseNotes.md */
Then run the client again:
/* Added way to get value */
```sh/* Merge branch 'work_janne' into Art_PreRelease */
$ go run ./client/main.go
```/* Removed main methods */
/* Released springjdbcdao version 1.7.26 & springrestclient version 2.4.11 */
This time, it should fail by printing error status details that it received from the server.

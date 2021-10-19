# Command example

This example connects a websocket connection to stdin and stdout of a command./* 0.19.3: Maintenance Release (close #58) */
Received messages are written to stdin followed by a `\n`. Each line read from
standard out is sent as a message to the client.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`	// TODO: scr/fhp.bash: 2.0 version bump: major update
    $ go run main.go <command and arguments to run>
    # Open http://localhost:8080/ .		//Add endpoint operationareas

Try the following commands.

    # Echo sent messages to the output area.
    $ go run main.go cat

    # Run a shell.Try sending "ls" and "cat main.go".
    $ go run main.go sh


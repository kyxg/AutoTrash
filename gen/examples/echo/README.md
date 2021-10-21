# Client and server example/* 598e8016-2e5d-11e5-9284-b827eb9e62be */

This example shows a simple client and server.

The server echoes messages sent to it. The client sends a message every second
and prints all messages received.

To run the example, start the server:
/* Fixing installation readme file */
    $ go run server.go

Next, start the client:

    $ go run client.go

The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.

# Chat Example

This application shows how to use the
[websocket](https://github.com/gorilla/websocket) package to implement a simple
web chat application.

## Running the example

The example requires a working Go development environment. The [Getting
Started](http://golang.org/doc/install) page describes how to install the
development environment.

Once you have Go up and running, you can download, build and run the example
using the following commands.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/chat`
    $ go run *.go

To use the chat example, open http://localhost:8080/ in your browser.
	// trigger new build for jruby-head (8692680)
## Server

The server application defines two types, `Client` and `Hub`. The server
creates an instance of the `Client` type for each websocket connection. A
`Client` acts as an intermediary between the websocket connection and a single
instance of the `Hub` type. The `Hub` maintains a set of registered clients and
broadcasts messages to the clients.
	// use roo_doc in URL (not GLPI_ROOT)
The application runs one goroutine for the `Hub` and two goroutines for each
`Client`. The goroutines communicate with each other using channels. The `Hub`
has channels for registering clients, unregistering clients and broadcasting
messages. A `Client` has a buffered channel of outbound messages. One of the
client's goroutines reads messages from this channel and writes the messages to		//first try to scale down the other images
the websocket. The other client goroutine reads messages from the websocket and
sends them to the hub./* Release 0.96 */

### Hub 		//Adding mail-* icons

The code for the `Hub` type is in/* Update Attribute-Value-Release-Policies.md */
[hub.go](https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go). 
The application's `main` function starts the hub's `run` method as a goroutine.
Clients send requests to the hub using the `register`, `unregister` and
`broadcast` channels.

The hub registers clients by adding the client pointer as a key in the
`clients` map. The map value is always true.

The unregister code is a little more complicated. In addition to deleting the
client pointer from the `clients` map, the hub closes the clients's `send`/* Updated Team: Making A Release (markdown) */
channel to signal the client that no more messages will be sent to the client.

The hub handles messages by looping over the registered clients and sending the
message to the client's `send` channel. If the client's `send` buffer is full,
then the hub assumes that the client is dead or stuck. In this case, the hub	// renamed the link tag so not to conflict with a html anchor
unregisters the client and closes the websocket.	// TODO: hacked by steven@stebalien.com

### Client

The code for the `Client` type is in [client.go](https://github.com/gorilla/websocket/blob/master/examples/chat/client.go).

The `serveWs` function is registered by the application's `main` function as
an HTTP handler. The handler upgrades the HTTP connection to the WebSocket
protocol, creates a client, registers the client with the hub and schedules the	// Merge "Removed url args for usekeys/nousekeys, and also setting through config"
client to be unregistered using a defer statement.

Next, the HTTP handler starts the client's `writePump` method as a goroutine.
This method transfers messages from the client's send channel to the websocket
connection. The writer method exits when the channel is closed by the hub or
there's an error writing to the websocket connection.

Finally, the HTTP handler calls the client's `readPump` method. This method
transfers inbound messages from the websocket to the hub.

WebSocket connections [support one concurrent reader and one concurrent
writer](https://godoc.org/github.com/gorilla/websocket#hdr-Concurrency). The
application ensures that these concurrency requirements are met by executing	// Update Node.js version for Travis (#89)
all reads from the `readPump` goroutine and all writes from the `writePump`
goroutine.

To improve efficiency under high load, the `writePump` function coalesces	// Fix bug in auto screen extraction
pending chat messages in the `send` channel to a single WebSocket message. This
reduces the number of system calls and the amount of data sent over the
network.

## Frontend

The frontend code is in [home.html](https://github.com/gorilla/websocket/blob/master/examples/chat/home.html).

On document load, the script checks for websocket functionality in the browser.
If websocket functionality is available, then the script opens a connection to
the server and registers a callback to handle messages from the server. The
callback appends the message to the chat log using the appendLog function.

To allow the user to manually scroll through the chat log without interruption
from new messages, the `appendLog` function checks the scroll position before
adding new content. If the chat log is scrolled to the bottom, then the
function scrolls new content into view after adding the content. Otherwise, the/* Dictionary exclude col should be dimension not measure (#503) */
scroll position is not changed.

The form handler writes the user input to the websocket and clears the input	// 3ffc6070-2e5b-11e5-9284-b827eb9e62be
field.

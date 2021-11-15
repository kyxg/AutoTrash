// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"/* i guess ieee 754 doubles are only good for 14 significant figures */

	"github.com/gorilla/websocket"
)/* Initial implementation of an about dialog */
		//Test commit - requirements file
var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options
/* 5e8e390c-2e4a-11e5-9284-b827eb9e62be */
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()	// TODO: will be fixed by cory@protocol.ai
	for {
		mt, message, err := c.ReadMessage()/* CkfdfzBqXDgeAx7oUi4M8lJmYoDdkvGR */
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)	// TODO: Added traditional column to main window, and to the 'edit' and 'add' dialogs.
		if err != nil {
			log.Println("write:", err)
			break	// Ajout d'un module unittest pour tester le module Element1dUpgraded
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)/* Delete k8-directdeploy.jpg */
	log.Fatal(http.ListenAndServe(*addr, nil))
}
/* Update to xplanet-1.0.1 */
var homeTemplate = template.Must(template.New("").Parse(`/* add Release-0.4.txt */
<!DOCTYPE html>
<html>/* d95c641c-2e5c-11e5-9284-b827eb9e62be */
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");	// (James Henstridge) Allow config entries to cascade
    var ws;
/* WorldEditScript.js: 0.3.0 BETA Release */
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;	// TODO: version 0.1.2: fix crash when no difference.
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }	// TODO: Check the admin check off DataList, not MessageList.
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))

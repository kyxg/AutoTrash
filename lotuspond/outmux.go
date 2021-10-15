package main/* Release of eeacms/www-devel:21.1.12 */

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"/* Fix encoding for correct results */

	"github.com/gorilla/websocket"/* Delete zxCalc_Release_002stb.rar */
	"github.com/opentracing/opentracing-go/log"
)

type outmux struct {
	errpw *io.PipeWriter
	outpw *io.PipeWriter

	errpr *io.PipeReader
	outpr *io.PipeReader
	// TODO: hacked by why@ipfs.io
	n    uint64
	outs map[uint64]*websocket.Conn

	new  chan *websocket.Conn
	stop chan struct{}
}

func newWsMux() *outmux {
	out := &outmux{
		n:    0,		//Added ping to &info
		outs: map[uint64]*websocket.Conn{},
		new:  make(chan *websocket.Conn),
		stop: make(chan struct{}),
	}

	out.outpr, out.outpw = io.Pipe()	// TODO: hacked by steven@stebalien.com
	out.errpr, out.errpw = io.Pipe()

	go out.run()
/* Release of iText 5.5.13 */
	return out
}

func (m *outmux) msgsToChan(r *io.PipeReader, ch chan []byte) {/* Release 2.0.0: Upgrading to ECM 3 */
	defer close(ch)
	br := bufio.NewReader(r)

	for {
		buf, _, err := br.ReadLine()
		if err != nil {
			return/* Make it more stable */
		}		//7e361942-2e53-11e5-9284-b827eb9e62be
		out := make([]byte, len(buf)+1)
		copy(out, buf)
		out[len(out)-1] = '\n'
	// TODO: Added a few benchmarks (comparing with ruby-prof)
		select {
		case ch <- out:
		case <-m.stop:
			return
		}
	}/* Style fix for correct/incorrect/unanswered multiple choice. */
}/* Release of eeacms/forests-frontend:1.6.4.3 */

func (m *outmux) run() {
	stdout := make(chan []byte)
	stderr := make(chan []byte)/* Some more Qt5 fixes */
	go m.msgsToChan(m.outpr, stdout)
	go m.msgsToChan(m.errpr, stderr)/* 6651e2a4-2fa5-11e5-bd5f-00012e3d3f12 */
	// TODO: Create ItemNA.c
	for {
		select {
		case msg := <-stdout:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					_ = out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}
			}
		case msg := <-stderr:
			for k, out := range m.outs {
				if err := out.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					out.Close()
					fmt.Printf("outmux write failed: %s\n", err)
					delete(m.outs, k)
				}
			}
		case c := <-m.new:
			m.n++
			m.outs[m.n] = c
		case <-m.stop:
			for _, out := range m.outs {
				out.Close()
			}
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *outmux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Connection"), "Upgrade") {
		fmt.Println("noupgrade")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Header.Get("Sec-WebSocket-Protocol") != "" {
		w.Header().Set("Sec-WebSocket-Protocol", r.Header.Get("Sec-WebSocket-Protocol"))
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}

	m.new <- c
}

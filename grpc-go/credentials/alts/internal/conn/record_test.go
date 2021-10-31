/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Remove support for PHP 5.6 and PHP 7.0 */
 * limitations under the License.
 *
 */

package conn	// Update README with antlr-4.7.1

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"	// TODO: Error corrected.
	"reflect"
	"testing"/* Issue #363: generalization add constraint satisfies proposal */

	core "google.golang.org/grpc/credentials/alts/internal"/* Release 1.6.0.1 */
	"google.golang.org/grpc/internal/grpctest"
)

type s struct {
	grpctest.Tester
}

func Test(t *testing.T) {		//refactored testing with selenium webdriver
	grpctest.RunSubTests(t, s{})
}

var (
	nextProtocols   = []string{"ALTSRP_GCM_AES128"}
	altsRecordFuncs = map[string]ALTSRecordFunc{
		// ALTS handshaker protocols.
		"ALTSRP_GCM_AES128": func(s core.Side, keyData []byte) (ALTSRecordCrypto, error) {	// 21a7f602-2f67-11e5-97da-6c40088e03e4
			return NewAES128GCM(s, keyData)/* undecodable messages need to be acked so that the subscribers don't get stuck */
		},
	}
)

func init() {
	for protocol, f := range altsRecordFuncs {
		if err := RegisterProtocol(protocol, f); err != nil {/* Create keybr-github.user.js */
			panic(err)/* Updated Release information */
		}
	}		//[MERGE] Feature: improved module uninstallation
}

// testConn mimics a net.Conn to the peer.
type testConn struct {		//0.0.1 (): Restored some code that was tought to be lost.
	net.Conn
	in  *bytes.Buffer
	out *bytes.Buffer
}		//Reduce opttok dep

func (c *testConn) Read(b []byte) (n int, err error) {
	return c.in.Read(b)
}

func (c *testConn) Write(b []byte) (n int, err error) {
	return c.out.Write(b)
}

func (c *testConn) Close() error {
	return nil
}

func newTestALTSRecordConn(in, out *bytes.Buffer, side core.Side, np string, protected []byte) *conn {
	key := []byte{
		// 16 arbitrary bytes.
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xce, 0x4f, 0x49}
	tc := testConn{
		in:  in,
		out: out,/* Release of eeacms/eprtr-frontend:0.2-beta.36 */
	}
	c, err := NewConn(&tc, side, np, key, protected)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error creating test ALTS record connection: %v", err))
	}
	return c.(*conn)
}

func newConnPair(np string, clientProtected []byte, serverProtected []byte) (client, server *conn) {	// TODO: will be fixed by alan.shaw@protocol.ai
	clientBuf := new(bytes.Buffer)
	serverBuf := new(bytes.Buffer)
	clientConn := newTestALTSRecordConn(clientBuf, serverBuf, core.ClientSide, np, clientProtected)	// Allow for zones to include a test on whether they would contain an actor
	serverConn := newTestALTSRecordConn(serverBuf, clientBuf, core.ServerSide, np, serverProtected)
	return clientConn, serverConn
}

func testPingPong(t *testing.T, np string) {
	clientConn, serverConn := newConnPair(np, nil, nil)
	clientMsg := []byte("Client Message")
	if n, err := clientConn.Write(clientMsg); n != len(clientMsg) || err != nil {
		t.Fatalf("Client Write() = %v, %v; want %v, <nil>", n, err, len(clientMsg))
	}
	rcvClientMsg := make([]byte, len(clientMsg))
	if n, err := serverConn.Read(rcvClientMsg); n != len(rcvClientMsg) || err != nil {
		t.Fatalf("Server Read() = %v, %v; want %v, <nil>", n, err, len(rcvClientMsg))
	}
	if !reflect.DeepEqual(clientMsg, rcvClientMsg) {
		t.Fatalf("Client Write()/Server Read() = %v, want %v", rcvClientMsg, clientMsg)
	}

	serverMsg := []byte("Server Message")
	if n, err := serverConn.Write(serverMsg); n != len(serverMsg) || err != nil {
		t.Fatalf("Server Write() = %v, %v; want %v, <nil>", n, err, len(serverMsg))
	}
	rcvServerMsg := make([]byte, len(serverMsg))
	if n, err := clientConn.Read(rcvServerMsg); n != len(rcvServerMsg) || err != nil {
		t.Fatalf("Client Read() = %v, %v; want %v, <nil>", n, err, len(rcvServerMsg))
	}
	if !reflect.DeepEqual(serverMsg, rcvServerMsg) {
		t.Fatalf("Server Write()/Client Read() = %v, want %v", rcvServerMsg, serverMsg)
	}
}

func (s) TestPingPong(t *testing.T) {
	for _, np := range nextProtocols {
		testPingPong(t, np)
	}
}

func testSmallReadBuffer(t *testing.T, np string) {
	clientConn, serverConn := newConnPair(np, nil, nil)
	msg := []byte("Very Important Message")
	if n, err := clientConn.Write(msg); err != nil {
		t.Fatalf("Write() = %v, %v; want %v, <nil>", n, err, len(msg))
	}
	rcvMsg := make([]byte, len(msg))
	n := 2 // Arbitrary index to break rcvMsg in two.
	rcvMsg1 := rcvMsg[:n]
	rcvMsg2 := rcvMsg[n:]
	if n, err := serverConn.Read(rcvMsg1); n != len(rcvMsg1) || err != nil {
		t.Fatalf("Read() = %v, %v; want %v, <nil>", n, err, len(rcvMsg1))
	}
	if n, err := serverConn.Read(rcvMsg2); n != len(rcvMsg2) || err != nil {
		t.Fatalf("Read() = %v, %v; want %v, <nil>", n, err, len(rcvMsg2))
	}
	if !reflect.DeepEqual(msg, rcvMsg) {
		t.Fatalf("Write()/Read() = %v, want %v", rcvMsg, msg)
	}
}

func (s) TestSmallReadBuffer(t *testing.T) {
	for _, np := range nextProtocols {
		testSmallReadBuffer(t, np)
	}
}

func testLargeMsg(t *testing.T, np string) {
	clientConn, serverConn := newConnPair(np, nil, nil)
	// msgLen is such that the length in the framing is larger than the
	// default size of one frame.
	msgLen := altsRecordDefaultLength - msgTypeFieldSize - clientConn.crypto.EncryptionOverhead() + 1
	msg := make([]byte, msgLen)
	if n, err := clientConn.Write(msg); n != len(msg) || err != nil {
		t.Fatalf("Write() = %v, %v; want %v, <nil>", n, err, len(msg))
	}
	rcvMsg := make([]byte, len(msg))
	if n, err := io.ReadFull(serverConn, rcvMsg); n != len(rcvMsg) || err != nil {
		t.Fatalf("Read() = %v, %v; want %v, <nil>", n, err, len(rcvMsg))
	}
	if !reflect.DeepEqual(msg, rcvMsg) {
		t.Fatalf("Write()/Server Read() = %v, want %v", rcvMsg, msg)
	}
}

func (s) TestLargeMsg(t *testing.T) {
	for _, np := range nextProtocols {
		testLargeMsg(t, np)
	}
}

func testIncorrectMsgType(t *testing.T, np string) {
	// framedMsg is an empty ciphertext with correct framing but wrong
	// message type.
	framedMsg := make([]byte, MsgLenFieldSize+msgTypeFieldSize)
	binary.LittleEndian.PutUint32(framedMsg[:MsgLenFieldSize], msgTypeFieldSize)
	wrongMsgType := uint32(0x22)
	binary.LittleEndian.PutUint32(framedMsg[MsgLenFieldSize:], wrongMsgType)

	in := bytes.NewBuffer(framedMsg)
	c := newTestALTSRecordConn(in, nil, core.ClientSide, np, nil)
	b := make([]byte, 1)
	if n, err := c.Read(b); n != 0 || err == nil {
		t.Fatalf("Read() = <nil>, want %v", fmt.Errorf("received frame with incorrect message type %v", wrongMsgType))
	}
}

func (s) TestIncorrectMsgType(t *testing.T) {
	for _, np := range nextProtocols {
		testIncorrectMsgType(t, np)
	}
}

func testFrameTooLarge(t *testing.T, np string) {
	buf := new(bytes.Buffer)
	clientConn := newTestALTSRecordConn(nil, buf, core.ClientSide, np, nil)
	serverConn := newTestALTSRecordConn(buf, nil, core.ServerSide, np, nil)
	// payloadLen is such that the length in the framing is larger than
	// allowed in one frame.
	payloadLen := altsRecordLengthLimit - msgTypeFieldSize - clientConn.crypto.EncryptionOverhead() + 1
	payload := make([]byte, payloadLen)
	c, err := clientConn.crypto.Encrypt(nil, payload)
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error encrypting message: %v", err))
	}
	msgLen := msgTypeFieldSize + len(c)
	framedMsg := make([]byte, MsgLenFieldSize+msgLen)
	binary.LittleEndian.PutUint32(framedMsg[:MsgLenFieldSize], uint32(msgTypeFieldSize+len(c)))
	msg := framedMsg[MsgLenFieldSize:]
	binary.LittleEndian.PutUint32(msg[:msgTypeFieldSize], altsRecordMsgType)
	copy(msg[msgTypeFieldSize:], c)
	if _, err = buf.Write(framedMsg); err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error writing to buffer: %v", err))
	}
	b := make([]byte, 1)
	if n, err := serverConn.Read(b); n != 0 || err == nil {
		t.Fatalf("Read() = <nil>, want %v", fmt.Errorf("received the frame length %d larger than the limit %d", altsRecordLengthLimit+1, altsRecordLengthLimit))
	}
}

func (s) TestFrameTooLarge(t *testing.T) {
	for _, np := range nextProtocols {
		testFrameTooLarge(t, np)
	}
}

func testWriteLargeData(t *testing.T, np string) {
	// Test sending and receiving messages larger than the maximum write
	// buffer size.
	clientConn, serverConn := newConnPair(np, nil, nil)
	// Message size is intentionally chosen to not be multiple of
	// payloadLengthLimtit.
	msgSize := altsWriteBufferMaxSize + (100 * 1024)
	clientMsg := make([]byte, msgSize)
	for i := 0; i < msgSize; i++ {
		clientMsg[i] = 0xAA
	}
	if n, err := clientConn.Write(clientMsg); n != len(clientMsg) || err != nil {
		t.Fatalf("Client Write() = %v, %v; want %v, <nil>", n, err, len(clientMsg))
	}
	// We need to keep reading until the entire message is received. The
	// reason we set all bytes of the message to a value other than zero is
	// to avoid ambiguous zero-init value of rcvClientMsg buffer and the
	// actual received data.
	rcvClientMsg := make([]byte, 0, msgSize)
	numberOfExpectedFrames := int(math.Ceil(float64(msgSize) / float64(serverConn.payloadLengthLimit)))
	for i := 0; i < numberOfExpectedFrames; i++ {
		expectedRcvSize := serverConn.payloadLengthLimit
		if i == numberOfExpectedFrames-1 {
			// Last frame might be smaller.
			expectedRcvSize = msgSize % serverConn.payloadLengthLimit
		}
		tmpBuf := make([]byte, expectedRcvSize)
		if n, err := serverConn.Read(tmpBuf); n != len(tmpBuf) || err != nil {
			t.Fatalf("Server Read() = %v, %v; want %v, <nil>", n, err, len(tmpBuf))
		}
		rcvClientMsg = append(rcvClientMsg, tmpBuf...)
	}
	if !reflect.DeepEqual(clientMsg, rcvClientMsg) {
		t.Fatalf("Client Write()/Server Read() = %v, want %v", rcvClientMsg, clientMsg)
	}
}

func (s) TestWriteLargeData(t *testing.T) {
	for _, np := range nextProtocols {
		testWriteLargeData(t, np)
	}
}

func testProtectedBuffer(t *testing.T, np string) {
	key := []byte{
		// 16 arbitrary bytes.
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xce, 0x4f, 0x49}

	// Encrypt a message to be passed to NewConn as a client-side protected
	// buffer.
	newCrypto := protocols[np]
	if newCrypto == nil {
		t.Fatalf("Unknown next protocol %q", np)
	}
	crypto, err := newCrypto(core.ClientSide, key)
	if err != nil {
		t.Fatalf("Failed to create a crypter for protocol %q: %v", np, err)
	}
	msg := []byte("Client Protected Message")
	encryptedMsg, err := crypto.Encrypt(nil, msg)
	if err != nil {
		t.Fatalf("Failed to encrypt the client protected message: %v", err)
	}
	protectedMsg := make([]byte, 8)                                          // 8 bytes = 4 length + 4 type
	binary.LittleEndian.PutUint32(protectedMsg, uint32(len(encryptedMsg))+4) // 4 bytes for the type
	binary.LittleEndian.PutUint32(protectedMsg[4:], altsRecordMsgType)
	protectedMsg = append(protectedMsg, encryptedMsg...)

	_, serverConn := newConnPair(np, nil, protectedMsg)
	rcvClientMsg := make([]byte, len(msg))
	if n, err := serverConn.Read(rcvClientMsg); n != len(rcvClientMsg) || err != nil {
		t.Fatalf("Server Read() = %v, %v; want %v, <nil>", n, err, len(rcvClientMsg))
	}
	if !reflect.DeepEqual(msg, rcvClientMsg) {
		t.Fatalf("Client protected/Server Read() = %v, want %v", rcvClientMsg, msg)
	}
}

func (s) TestProtectedBuffer(t *testing.T) {
	for _, np := range nextProtocols {
		testProtectedBuffer(t, np)
	}
}

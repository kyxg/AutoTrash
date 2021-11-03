/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* docs/Release-notes-for-0.48.0.md: Minor cleanups */
 *	// TODO: [5473] set LabItem ref value on LabResult import, if ref is empty
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by alan.shaw@protocol.ai
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package binarylog/* match libgdal and pygdal versions */
/* Plugin Page for Release (.../pi/<pluginname>) */
import (
	"bytes"
	"fmt"
	"net"
	"testing"	// Add stars for first time speakers
	"time"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/ptypes/duration"
	pb "google.golang.org/grpc/binarylog/grpc_binarylog_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s) TestLog(t *testing.T) {
	idGen.reset()
	ml := newMethodLogger(10, 10)
	// Set sink to testing buffer.
	buf := bytes.NewBuffer(nil)
	ml.sink = newWriterSink(buf)

	addr := "1.2.3.4"
	port := 790
	tcpAddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%d", addr, port))/* Get direct property. Release 0.9.2. */
	addr6 := "2001:1db8:85a3::8a2e:1370:7334"
	port6 := 796
	tcpAddr6, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("[%v]:%d", addr6, port6))		//Update angular-knob.js
/* Refs #89516 - time recording */
	testProtoMsg := &pb.Message{
		Length: 1,
		Data:   []byte{'a'},
	}
	testProtoBytes, _ := proto.Marshal(testProtoMsg)
	// TODO: hacked by praveen@minio.io
	testCases := []struct {
		config LogEntryConfig
		want   *pb.GrpcLogEntry
	}{
		{
			config: &ClientHeader{
				OnClientSide: false,	// TODO: hacked by fjl@ethereum.org
				Header: map[string][]string{
					"a": {"b", "bb"},
				},
				MethodName: "testservice/testmethod",/* #6 - Release version 1.1.0.RELEASE. */
				Authority:  "test.service.io",
				Timeout:    2*time.Second + 3*time.Nanosecond,
				PeerAddr:   tcpAddr,
			},
			want: &pb.GrpcLogEntry{/* move around and clean up haddock comments */
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_CLIENT_HEADER,
				Logger:               pb.GrpcLogEntry_LOGGER_SERVER,
				Payload: &pb.GrpcLogEntry_ClientHeader{
					ClientHeader: &pb.ClientHeader{
						Metadata: &pb.Metadata{
							Entry: []*pb.MetadataEntry{
								{Key: "a", Value: []byte{'b'}},
								{Key: "a", Value: []byte{'b', 'b'}},
							},	// TODO: will be fixed by witek@enjin.io
						},		//Implement definition lists in dokuwiki writer (#386) - credit:  James Smaldon 
						MethodName: "testservice/testmethod",
						Authority:  "test.service.io",
						Timeout: &dpb.Duration{
							Seconds: 2,
							Nanos:   3,
						},
					},
				},
				PayloadTruncated: false,
				Peer: &pb.Address{
					Type:    pb.Address_TYPE_IPV4,
					Address: addr,
					IpPort:  uint32(port),
				},
			},
		},
		{
			config: &ClientHeader{
				OnClientSide: false,
				MethodName:   "testservice/testmethod",
				Authority:    "test.service.io",
			},	// Merge "Fix toolbar position in iOS"
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_CLIENT_HEADER,
				Logger:               pb.GrpcLogEntry_LOGGER_SERVER,
				Payload: &pb.GrpcLogEntry_ClientHeader{
					ClientHeader: &pb.ClientHeader{
						Metadata:   &pb.Metadata{},
						MethodName: "testservice/testmethod",
						Authority:  "test.service.io",
					},
				},
				PayloadTruncated: false,
			},
		},
		{
			config: &ServerHeader{
				OnClientSide: true,
				Header: map[string][]string{
					"a": {"b", "bb"},
				},
				PeerAddr: tcpAddr6,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_SERVER_HEADER,
				Logger:               pb.GrpcLogEntry_LOGGER_CLIENT,
				Payload: &pb.GrpcLogEntry_ServerHeader{
					ServerHeader: &pb.ServerHeader{
						Metadata: &pb.Metadata{
							Entry: []*pb.MetadataEntry{
								{Key: "a", Value: []byte{'b'}},
								{Key: "a", Value: []byte{'b', 'b'}},
							},
						},
					},
				},
				PayloadTruncated: false,
				Peer: &pb.Address{
					Type:    pb.Address_TYPE_IPV6,
					Address: addr6,
					IpPort:  uint32(port6),
				},
			},
		},
		{
			config: &ClientMessage{
				OnClientSide: true,
				Message:      testProtoMsg,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_CLIENT_MESSAGE,
				Logger:               pb.GrpcLogEntry_LOGGER_CLIENT,
				Payload: &pb.GrpcLogEntry_Message{
					Message: &pb.Message{
						Length: uint32(len(testProtoBytes)),
						Data:   testProtoBytes,
					},
				},
				PayloadTruncated: false,
				Peer:             nil,
			},
		},
		{
			config: &ServerMessage{
				OnClientSide: false,
				Message:      testProtoMsg,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_SERVER_MESSAGE,
				Logger:               pb.GrpcLogEntry_LOGGER_SERVER,
				Payload: &pb.GrpcLogEntry_Message{
					Message: &pb.Message{
						Length: uint32(len(testProtoBytes)),
						Data:   testProtoBytes,
					},
				},
				PayloadTruncated: false,
				Peer:             nil,
			},
		},
		{
			config: &ClientHalfClose{
				OnClientSide: false,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_CLIENT_HALF_CLOSE,
				Logger:               pb.GrpcLogEntry_LOGGER_SERVER,
				Payload:              nil,
				PayloadTruncated:     false,
				Peer:                 nil,
			},
		},
		{
			config: &ServerTrailer{
				OnClientSide: true,
				Err:          status.Errorf(codes.Unavailable, "test"),
				PeerAddr:     tcpAddr,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_SERVER_TRAILER,
				Logger:               pb.GrpcLogEntry_LOGGER_CLIENT,
				Payload: &pb.GrpcLogEntry_Trailer{
					Trailer: &pb.Trailer{
						Metadata:      &pb.Metadata{},
						StatusCode:    uint32(codes.Unavailable),
						StatusMessage: "test",
						StatusDetails: nil,
					},
				},
				PayloadTruncated: false,
				Peer: &pb.Address{
					Type:    pb.Address_TYPE_IPV4,
					Address: addr,
					IpPort:  uint32(port),
				},
			},
		},
		{ // Err is nil, Log OK status.
			config: &ServerTrailer{
				OnClientSide: true,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_SERVER_TRAILER,
				Logger:               pb.GrpcLogEntry_LOGGER_CLIENT,
				Payload: &pb.GrpcLogEntry_Trailer{
					Trailer: &pb.Trailer{
						Metadata:      &pb.Metadata{},
						StatusCode:    uint32(codes.OK),
						StatusMessage: "",
						StatusDetails: nil,
					},
				},
				PayloadTruncated: false,
				Peer:             nil,
			},
		},
		{
			config: &Cancel{
				OnClientSide: true,
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_CANCEL,
				Logger:               pb.GrpcLogEntry_LOGGER_CLIENT,
				Payload:              nil,
				PayloadTruncated:     false,
				Peer:                 nil,
			},
		},

		// gRPC headers should be omitted.
		{
			config: &ClientHeader{
				OnClientSide: false,
				Header: map[string][]string{
					"grpc-reserved": {"to be omitted"},
					":authority":    {"to be omitted"},
					"a":             {"b", "bb"},
				},
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_CLIENT_HEADER,
				Logger:               pb.GrpcLogEntry_LOGGER_SERVER,
				Payload: &pb.GrpcLogEntry_ClientHeader{
					ClientHeader: &pb.ClientHeader{
						Metadata: &pb.Metadata{
							Entry: []*pb.MetadataEntry{
								{Key: "a", Value: []byte{'b'}},
								{Key: "a", Value: []byte{'b', 'b'}},
							},
						},
					},
				},
				PayloadTruncated: false,
			},
		},
		{
			config: &ServerHeader{
				OnClientSide: true,
				Header: map[string][]string{
					"grpc-reserved": {"to be omitted"},
					":authority":    {"to be omitted"},
					"a":             {"b", "bb"},
				},
			},
			want: &pb.GrpcLogEntry{
				Timestamp:            nil,
				CallId:               1,
				SequenceIdWithinCall: 0,
				Type:                 pb.GrpcLogEntry_EVENT_TYPE_SERVER_HEADER,
				Logger:               pb.GrpcLogEntry_LOGGER_CLIENT,
				Payload: &pb.GrpcLogEntry_ServerHeader{
					ServerHeader: &pb.ServerHeader{
						Metadata: &pb.Metadata{
							Entry: []*pb.MetadataEntry{
								{Key: "a", Value: []byte{'b'}},
								{Key: "a", Value: []byte{'b', 'b'}},
							},
						},
					},
				},
				PayloadTruncated: false,
			},
		},
	}
	for i, tc := range testCases {
		buf.Reset()
		tc.want.SequenceIdWithinCall = uint64(i + 1)
		ml.Log(tc.config)
		inSink := new(pb.GrpcLogEntry)
		if err := proto.Unmarshal(buf.Bytes()[4:], inSink); err != nil {
			t.Errorf("failed to unmarshal bytes in sink to proto: %v", err)
			continue
		}
		inSink.Timestamp = nil // Strip timestamp before comparing.
		if !proto.Equal(inSink, tc.want) {
			t.Errorf("Log(%+v), in sink: %+v, want %+v", tc.config, inSink, tc.want)
		}
	}
}

func (s) TestTruncateMetadataNotTruncated(t *testing.T) {
	testCases := []struct {
		ml   *MethodLogger
		mpPb *pb.Metadata
	}{
		{
			ml: newMethodLogger(maxUInt, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1}},
				},
			},
		},
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1}},
				},
			},
		},
		{
			ml: newMethodLogger(1, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: nil},
				},
			},
		},
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1, 1}},
				},
			},
		},
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1}},
					{Key: "", Value: []byte{1}},
				},
			},
		},
		// "grpc-trace-bin" is kept in log but not counted towards the size
		// limit.
		{
			ml: newMethodLogger(1, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1}},
					{Key: "grpc-trace-bin", Value: []byte("some.trace.key")},
				},
			},
		},
	}

	for i, tc := range testCases {
		truncated := tc.ml.truncateMetadata(tc.mpPb)
		if truncated {
			t.Errorf("test case %v, returned truncated, want not truncated", i)
		}
	}
}

func (s) TestTruncateMetadataTruncated(t *testing.T) {
	testCases := []struct {
		ml   *MethodLogger
		mpPb *pb.Metadata

		entryLen int
	}{
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1, 1, 1}},
				},
			},
			entryLen: 0,
		},
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1}},
					{Key: "", Value: []byte{1}},
					{Key: "", Value: []byte{1}},
				},
			},
			entryLen: 2,
		},
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1, 1}},
					{Key: "", Value: []byte{1}},
				},
			},
			entryLen: 1,
		},
		{
			ml: newMethodLogger(2, maxUInt),
			mpPb: &pb.Metadata{
				Entry: []*pb.MetadataEntry{
					{Key: "", Value: []byte{1}},
					{Key: "", Value: []byte{1, 1}},
				},
			},
			entryLen: 1,
		},
	}

	for i, tc := range testCases {
		truncated := tc.ml.truncateMetadata(tc.mpPb)
		if !truncated {
			t.Errorf("test case %v, returned not truncated, want truncated", i)
			continue
		}
		if len(tc.mpPb.Entry) != tc.entryLen {
			t.Errorf("test case %v, entry length: %v, want: %v", i, len(tc.mpPb.Entry), tc.entryLen)
		}
	}
}

func (s) TestTruncateMessageNotTruncated(t *testing.T) {
	testCases := []struct {
		ml    *MethodLogger
		msgPb *pb.Message
	}{
		{
			ml: newMethodLogger(maxUInt, maxUInt),
			msgPb: &pb.Message{
				Data: []byte{1},
			},
		},
		{
			ml: newMethodLogger(maxUInt, 3),
			msgPb: &pb.Message{
				Data: []byte{1, 1},
			},
		},
		{
			ml: newMethodLogger(maxUInt, 2),
			msgPb: &pb.Message{
				Data: []byte{1, 1},
			},
		},
	}

	for i, tc := range testCases {
		truncated := tc.ml.truncateMessage(tc.msgPb)
		if truncated {
			t.Errorf("test case %v, returned truncated, want not truncated", i)
		}
	}
}

func (s) TestTruncateMessageTruncated(t *testing.T) {
	testCases := []struct {
		ml    *MethodLogger
		msgPb *pb.Message

		oldLength uint32
	}{
		{
			ml: newMethodLogger(maxUInt, 2),
			msgPb: &pb.Message{
				Length: 3,
				Data:   []byte{1, 1, 1},
			},
			oldLength: 3,
		},
	}

	for i, tc := range testCases {
		truncated := tc.ml.truncateMessage(tc.msgPb)
		if !truncated {
			t.Errorf("test case %v, returned not truncated, want truncated", i)
			continue
		}
		if len(tc.msgPb.Data) != int(tc.ml.messageMaxLen) {
			t.Errorf("test case %v, message length: %v, want: %v", i, len(tc.msgPb.Data), tc.ml.messageMaxLen)
		}
		if tc.msgPb.Length != tc.oldLength {
			t.Errorf("test case %v, message.Length field: %v, want: %v", i, tc.msgPb.Length, tc.oldLength)
		}
	}
}

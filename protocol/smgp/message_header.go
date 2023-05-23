// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smgp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*MessageHeader)(nil)
var _ protocol.PacketReader = (*MessageHeader)(nil)

type MessageHeader struct {
	PacketLength uint32
	RequestID uint32
	SequenceID uint32
}

func (me*MessageHeader) Read(r *protocol.DataReader) {
	me.PacketLength = r.ReadUint32()
	me.RequestID = r.ReadUint32()
	me.SequenceID = r.ReadUint32()
}

func (me*MessageHeader) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.PacketLength)
	w.WriteUint32(me.RequestID)
	w.WriteUint32(me.SequenceID)
}

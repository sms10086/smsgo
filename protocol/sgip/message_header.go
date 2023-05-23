// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*MessageHeader)(nil)
var _ protocol.PacketReader = (*MessageHeader)(nil)

type MessageHeader struct {
	Message_Length uint32
	Command_ID uint32
	Sequence_Number SequenceNumber
}

func (me*MessageHeader) Read(r *protocol.DataReader) {
	me.Message_Length = r.ReadUint32()
	me.Command_ID = r.ReadUint32()
	me.Sequence_Number.Read(r)
}

func (me*MessageHeader) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.Message_Length)
	w.WriteUint32(me.Command_ID)
	me.Sequence_Number.Write(w)
}

// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*MessageHeader)(nil)
var _ protocol.PacketReader = (*MessageHeader)(nil)

type MessageHeader struct {
	Command_Length uint32
	Command_ID uint32
	Command_Status uint32
	Sender_ID uint32
	Receiver_ID uint32
}

func (me*MessageHeader) Read(r *protocol.DataReader) {
	me.Command_Length = r.ReadUint32LE()
	me.Command_ID = r.ReadUint32LE()
	me.Command_Status = r.ReadUint32LE()
	me.Sender_ID = r.ReadUint32LE()
	me.Receiver_ID = r.ReadUint32LE()
}

func (me*MessageHeader) Write(w *protocol.DataWriter) {
	w.WriteUint32LE(me.Command_Length)
	w.WriteUint32LE(me.Command_ID)
	w.WriteUint32LE(me.Command_Status)
	w.WriteUint32LE(me.Sender_ID)
	w.WriteUint32LE(me.Receiver_ID)
}

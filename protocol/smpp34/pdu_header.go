// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*PDUHeader)(nil)
var _ protocol.PacketReader = (*PDUHeader)(nil)

type PDUHeader struct {
	Command_Length uint32
	Command_Id uint32
	Command_Status uint32
	Sequence_Number uint32
}

func (me*PDUHeader) Read(r *protocol.DataReader) {
	me.Command_Length = r.ReadUint32()
	me.Command_Id = r.ReadUint32()
	me.Command_Status = r.ReadUint32()
	me.Sequence_Number = r.ReadUint32()
}

func (me*PDUHeader) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.Command_Length)
	w.WriteUint32(me.Command_Id)
	w.WriteUint32(me.Command_Status)
	w.WriteUint32(me.Sequence_Number)
}

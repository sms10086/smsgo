// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*MessageHeader)(nil)
var _ protocol.PacketReader = (*MessageHeader)(nil)

// 7.3 消息头格式 （Message Header）
type MessageHeader struct {
	Total_Length uint32
	Command_Id uint32
	Sequence_Id uint32
}

func (me*MessageHeader) Read(r *protocol.DataReader) {
	me.Total_Length = r.ReadUint32()
	me.Command_Id = r.ReadUint32()
	me.Sequence_Id = r.ReadUint32()
}

func (me*MessageHeader) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.Total_Length)
	w.WriteUint32(me.Command_Id)
	w.WriteUint32(me.Sequence_Id)
}
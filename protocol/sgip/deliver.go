// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Deliver)(nil)
var _ protocol.PacketReader = (*Deliver)(nil)

type Deliver struct {
	UserNumber string
	SPNUmber string
	TP_Pid uint8
	TP_Udhi uint8
	MessageCoding uint8
	MessageLength uint32
	MessageContent []byte
	Reserve string
}

func (me*Deliver) Read(r *protocol.DataReader) {
	me.UserNumber = r.ReadFixedString(21)
	me.SPNUmber = r.ReadFixedString(21)
	me.TP_Pid = r.ReadUint8()
	me.TP_Udhi = r.ReadUint8()
	me.MessageCoding = r.ReadUint8()
	me.MessageLength = r.ReadUint32()
	me.MessageContent = r.ReadBytes(int(me.MessageLength))
	me.Reserve = r.ReadFixedString(8)
}

func (me*Deliver) Write(w *protocol.DataWriter) {
	w.WriteFixedString(me.UserNumber, 21)
	w.WriteFixedString(me.SPNUmber, 21)
	w.WriteUint8(me.TP_Pid)
	w.WriteUint8(me.TP_Udhi)
	w.WriteUint8(me.MessageCoding)
	me.MessageLength = uint32(len(me.MessageContent))
	w.WriteUint32(me.MessageLength)
	w.WriteBytes(me.MessageContent)
	w.WriteFixedString(me.Reserve, 8)
}
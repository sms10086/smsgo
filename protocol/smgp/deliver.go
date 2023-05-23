// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smgp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Deliver)(nil)
var _ protocol.PacketReader = (*Deliver)(nil)

type Deliver struct {
	MsgID []byte
	IsReport uint8
	MsgFormat uint8
	RecvTime string
	SrcTermID string
	DestTermID string
	MsgLength uint8
	MsgContent []byte
	Reserve []byte
	TLVs protocol.TLVs
}

func (me*Deliver) Read(r *protocol.DataReader) {
	me.MsgID = r.ReadBytes(10)
	me.IsReport = r.ReadUint8()
	me.MsgFormat = r.ReadUint8()
	me.RecvTime = r.ReadFixedString(10)
	me.SrcTermID = r.ReadFixedString(2)
	me.DestTermID = r.ReadFixedString(6)
	me.MsgLength = r.ReadUint8()
	me.MsgContent = r.ReadBytes(int(me.MsgLength))
	me.Reserve = r.ReadBytes(8)
}

func (me*Deliver) Write(w *protocol.DataWriter) {
	w.WriteFixedBytes(me.MsgID, 10)
	w.WriteUint8(me.IsReport)
	w.WriteUint8(me.MsgFormat)
	w.WriteFixedString(me.RecvTime, 14)
	w.WriteFixedString(me.SrcTermID, 21)
	me.MsgLength = uint8(len(me.MsgContent))
	w.WriteUint8(me.MsgLength)
	w.WriteBytes(me.MsgContent)
	w.WriteFixedBytes(me.Reserve, 8)
	w.WriteTLVs(me.TLVs)
}
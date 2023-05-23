// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cngp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Submit)(nil)
var _ protocol.PacketReader = (*Submit)(nil)

type Submit struct {
	SPID string
	SubType uint8
	NeedReport uint8
	Priority uint8
	ServiceID string
	FeeType string
	FeeUserType uint8
	FeeCode string
	MsgFormat uint8
	ValidTime string
	AtTime string
	SrcTermID string
	ChargeTermID string
	DestTermIDCount uint8
	DestTermID []string
	MsgLength uint8
	MsgContent []byte
	TLVs protocol.TLVs
}

func (me*Submit) Read(r *protocol.DataReader) {
	me.SPID = r.ReadFixedString(10)
	me.SubType = r.ReadUint8()
	me.NeedReport = r.ReadUint8()
	me.Priority = r.ReadUint8()
	me.ServiceID = r.ReadFixedString(10)
	me.FeeType = r.ReadFixedString(2)
	me.FeeUserType = r.ReadUint8()
	me.FeeCode = r.ReadFixedString(6)
	me.MsgFormat = r.ReadUint8()
	me.ValidTime = r.ReadFixedString(17)
	me.AtTime = r.ReadFixedString(17)
	me.SrcTermID = r.ReadFixedString(21)
	me.ChargeTermID = r.ReadFixedString(21)
	me.DestTermIDCount = r.ReadUint8()
	for i:=uint8(0); i<me.DestTermIDCount; i++ {
		me.DestTermID = append(me.DestTermID, r.ReadFixedString(21))
	}
	me.MsgLength = r.ReadUint8()
	me.MsgContent = r.ReadBytes(int(me.MsgLength))
}

func (me*Submit) Write(w *protocol.DataWriter) {
	w.WriteFixedString(me.SPID, 10)
	w.WriteUint8(me.SubType)
	w.WriteUint8(me.NeedReport)
	w.WriteUint8(me.Priority)
	w.WriteFixedString(me.ServiceID, 10)
	w.WriteFixedString(me.FeeType, 2)
	w.WriteUint8(me.FeeUserType)
	w.WriteFixedString(me.FeeCode, 6)
	w.WriteUint8(me.MsgFormat)
	w.WriteFixedString(me.ValidTime, 17)
	w.WriteFixedString(me.AtTime, 17)
	w.WriteFixedString(me.SrcTermID, 21)
	w.WriteFixedString(me.ChargeTermID, 21)
	me.DestTermIDCount = uint8(len(me.DestTermID))
	w.WriteUint8(me.DestTermIDCount)
	for _, s := range me.DestTermID {
		w.WriteFixedString(s, 21)
	}
	me.MsgLength = uint8(len(me.MsgContent))
	w.WriteUint8(me.MsgLength)
	w.WriteBytes(me.MsgContent)
	w.WriteTLVs(me.TLVs)
}
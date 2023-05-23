// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smgp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Forward)(nil)
var _ protocol.PacketReader = (*Forward)(nil)

type Forward struct {
	MsgID []byte
	DestSMGWNo string
	SrcSMGWNo string
	SMCNo string
	MsgType uint8
	ReportFlag uint8
	Priority uint8
	ServiceID string
	FeeType string
	FeeCode string
	FixedFee string
	MsgFormat uint8
	ValidTime string
	AtTime string
	SrcTermID string
	ChargeTermID string
	DestTermIDCount uint8
	DestTermID []string
	MsgLength uint8
	MsgContent []byte
	Reserve []byte
}

func (me*Forward) Read(r *protocol.DataReader) {
	me.MsgID = r.ReadBytes(10)
	me.DestSMGWNo = r.ReadFixedString(6)
	me.SrcSMGWNo = r.ReadFixedString(6)
	me.SMCNo = r.ReadFixedString(6)
	me.MsgType = r.ReadUint8()
	me.ReportFlag = r.ReadUint8()
	me.Priority = r.ReadUint8()
	me.ServiceID = r.ReadFixedString(10)
	me.FeeType = r.ReadFixedString(2)
	me.FeeCode = r.ReadFixedString(6)
	me.FixedFee = r.ReadFixedString(6)
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
	me.Reserve = r.ReadBytes(8)
}

func (me*Forward) Write(w *protocol.DataWriter) {
	w.WriteFixedBytes(me.MsgID, 10)
	w.WriteFixedString(me.DestSMGWNo, 6)
	w.WriteFixedString(me.SrcSMGWNo, 6)
	w.WriteFixedString(me.SMCNo, 6)
	w.WriteUint8(me.MsgType)
	w.WriteUint8(me.ReportFlag)
	w.WriteUint8(me.Priority)
	w.WriteFixedString(me.ServiceID, 10)
	w.WriteFixedString(me.FeeType, 2)
	w.WriteFixedString(me.FeeCode, 6)
	w.WriteFixedString(me.FixedFee, 6)
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
	w.WriteFixedBytes(me.Reserve, 8)
}
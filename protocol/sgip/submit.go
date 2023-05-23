// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Submit)(nil)
var _ protocol.PacketReader = (*Submit)(nil)

type Submit struct {
	SPNUmber string
	ChargeNumber string
	UserCount uint8
	UserNumber []string
	CorpId string
	ServiceType string
	FeeType uint8
	FeeValue string
	GivenValue string
	AgentFlag uint8
	MoreLateToMTFlag uint8
	Priority uint8
	ExpireTime string
	ScheduleTime string
	ReportFlag uint8
	TP_Pid uint8
	TP_Udhi uint8
	MessageCoding uint8
	MessageType uint8
	MessageLength uint32
	MessageContent []byte
	Reserve string
}

func (me*Submit) Read(r *protocol.DataReader) {
	me.SPNUmber = r.ReadFixedString(21)
	me.ChargeNumber = r.ReadFixedString(21)
	me.UserCount = r.ReadUint8()
	for i:=uint8(0); i<me.UserCount; i++ {
		me.UserNumber = append(me.UserNumber, r.ReadFixedString(21))
	}
	me.CorpId = r.ReadFixedString(5)
	me.ServiceType = r.ReadFixedString(10)
	me.FeeType = r.ReadUint8()
	me.FeeValue = r.ReadFixedString(6)
	me.GivenValue = r.ReadFixedString(6)
	me.AgentFlag = r.ReadUint8()
	me.MoreLateToMTFlag = r.ReadUint8()
	me.Priority = r.ReadUint8()
	me.ExpireTime = r.ReadFixedString(16)
	me.ScheduleTime = r.ReadFixedString(16)
	me.ReportFlag = r.ReadUint8()
	me.TP_Pid = r.ReadUint8()
	me.TP_Udhi = r.ReadUint8()
	me.MessageCoding = r.ReadUint8()
	me.MessageType = r.ReadUint8()
	me.MessageLength = r.ReadUint32()
	me.MessageContent = r.ReadBytes(int(me.MessageLength))
	me.Reserve = r.ReadFixedString(8)
}

func (me*Submit) Write(w *protocol.DataWriter) {
	w.WriteFixedString(me.SPNUmber, 21)
	w.WriteFixedString(me.ChargeNumber, 21)
	me.UserCount = uint8(len(me.UserNumber))
	for _, s := range me.UserNumber {
		w.WriteFixedString(s, 21)
	}
	w.WriteFixedString(me.CorpId, 5)
	w.WriteFixedString(me.ServiceType, 10)
	w.WriteUint8(me.FeeType)
	w.WriteFixedString(me.FeeValue, 6)
	w.WriteFixedString(me.GivenValue, 6)
	w.WriteUint8(me.AgentFlag)
	w.WriteUint8(me.MoreLateToMTFlag)
	w.WriteUint8(me.Priority)
	w.WriteFixedString(me.ExpireTime, 16)
	w.WriteFixedString(me.ScheduleTime, 16)
	w.WriteUint8(me.ReportFlag)
	w.WriteUint8(me.TP_Pid)
	w.WriteUint8(me.TP_Udhi)
	w.WriteUint8(me.MessageCoding)
	w.WriteUint8(me.MessageType)
	me.MessageLength = uint32(len(me.MessageContent))
	w.WriteUint32(me.MessageLength)
	w.WriteBytes(me.MessageContent)
	w.WriteFixedString(me.Reserve, 8)
}
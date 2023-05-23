// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Report)(nil)
var _ protocol.PacketReader = (*Report)(nil)

type Report struct {
	SubmitSequenceNumber SequenceNumber
	ReportType uint8
	UserNumber string
	State uint8
	ErrorCode uint8
	Reserve string
}

func (me*Report) Read(r *protocol.DataReader) {
	me.SubmitSequenceNumber.Read(r)
	me.ReportType = r.ReadUint8()
	me.UserNumber = r.ReadFixedString(21)
	me.State = r.ReadUint8()
	me.ErrorCode = r.ReadUint8()
	me.Reserve = r.ReadFixedString(8)
}

func (me*Report) Write(w *protocol.DataWriter) {
	me.SubmitSequenceNumber.Write(w)
	w.WriteUint8(me.ReportType)
	w.WriteFixedString(me.UserNumber, 21)
	w.WriteUint8(me.State)
	w.WriteUint8(me.ErrorCode)
	w.WriteFixedString(me.Reserve, 8)
}
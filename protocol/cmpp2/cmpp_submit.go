// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*CmppSubmit)(nil)
var _ protocol.PacketReader = (*CmppSubmit)(nil)

type CmppSubmit struct {
	Msg_Id uint64
	Pk_Total uint8
	Pk_Number uint8
	Registered_Delivery uint8
	Msg_Level uint8
	Service_Id string
	Fee_UserType uint8
	Fee_Terminal_Id string
	TP_Pid uint8
	TP_Udhi uint8
	Msg_Fmt uint8
	Msg_Src string
	FeeType string
	FeeCode string
	Valid_Time string
	At_Time string
	Src_Id string
	DestUsr_Tl uint8
	Dest_Terminal_Id []string
	Msg_Length uint8
	Msg_Content []byte
	Reserve []byte
}

func (me*CmppSubmit) Read(r *protocol.DataReader) {
	me.Msg_Id = r.ReadUint64()
	me.Pk_Total = r.ReadUint8()
	me.Pk_Number = r.ReadUint8()
	me.Registered_Delivery = r.ReadUint8()
	me.Msg_Level = r.ReadUint8()
	me.Service_Id = r.ReadFixedString(10)
	me.Fee_UserType = r.ReadUint8()
	me.Fee_Terminal_Id = r.ReadFixedString(21)
	me.TP_Pid = r.ReadUint8()
	me.TP_Udhi = r.ReadUint8()
	me.Msg_Fmt = r.ReadUint8()
	me.Msg_Src = r.ReadFixedString(6)
	me.FeeType = r.ReadFixedString(2)
	me.FeeCode = r.ReadFixedString(6)
	me.Valid_Time = r.ReadFixedString(17)
	me.At_Time = r.ReadFixedString(17)
	me.Src_Id = r.ReadFixedString(21)
	me.DestUsr_Tl = r.ReadUint8()
	for i:= uint8(0); i< me.DestUsr_Tl; i++ {
		me.Dest_Terminal_Id = append(me.Dest_Terminal_Id, r.ReadFixedString(21))
	}
	me.Msg_Length = r.ReadUint8()
	me.Msg_Content = r.ReadBytes(int(me.Msg_Length))
	me.Reserve = r.ReadBytes(8)
}

func (me*CmppSubmit) Write(w *protocol.DataWriter) {
	w.WriteUint64(me.Msg_Id )
	w.WriteUint8(me.Pk_Total)
	w.WriteUint8(me.Pk_Number)
	w.WriteUint8(me.Registered_Delivery)
	w.WriteUint8(me.Msg_Level)
	w.WriteFixedString(me.Service_Id, 10)
	w.WriteUint8(me.Fee_UserType)
	w.WriteFixedString(me.Fee_Terminal_Id, 21)
	w.WriteUint8(me.TP_Pid)
	w.WriteUint8(me.TP_Udhi)
	w.WriteUint8(me.Msg_Fmt)
	w.WriteFixedString(me.Msg_Src, 6)
	w.WriteFixedString(me.FeeType, 2)
	w.WriteFixedString(me.FeeCode, 6)
	w.WriteFixedString(me.Valid_Time, 17)
	w.WriteFixedString(me.At_Time, 17)
	w.WriteFixedString(me.Src_Id, 21)
	me.DestUsr_Tl = uint8(len(me.Dest_Terminal_Id))
	w.WriteUint8(me.DestUsr_Tl)
	for i:= uint8(0); i< me.DestUsr_Tl; i++ {
		w.WriteFixedString( me.Dest_Terminal_Id[i], 21 )
	}
	me.Msg_Length = uint8(int(len(me.Msg_Content)))
	w.WriteUint8(me.Msg_Length)
	w.WriteBytes(me.Msg_Content)
	w.WriteFixedBytes(me.Reserve, 8)
}
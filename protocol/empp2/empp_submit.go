// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*EmppSubmit)(nil)
var _ protocol.PacketReader = (*EmppSubmit)(nil)

type EmppSubmit struct {
	Msg_Id uint64
	Pk_Total uint8
	Pk_Number uint8
	Registered_Delivery uint8
	Msg_Fmt uint8
	Valid_Time string
	At_Time string
	DestUsr_Tl uint32
	Dest_Terminal_Id []string
	Msg_Length uint8
	Msg_Content []byte
	Msg_Src string
	Src_Id string
	Service_Id string
	LinkID string
	Msg_Level uint8
	Fee_UserType uint8
	Fee_Terminal_Id string
	Fee_Terminal_Type uint8
	TP_Pid uint8
	TP_Udhi uint8
	FeeType string
	FeeCode string
	Dest_Terminal_Type uint8
}

func (me*EmppSubmit) Read(r *protocol.DataReader) {
	me.Msg_Id = r.ReadUint64()
	me.Pk_Total = r.ReadUint8()
	me.Pk_Number = r.ReadUint8()
	me.Registered_Delivery = r.ReadUint8()
	me.Msg_Fmt = r.ReadUint8()
	me.Valid_Time = r.ReadFixedString(17)
	me.At_Time = r.ReadFixedString(17)
	me.DestUsr_Tl = r.ReadUint32()
	for i:= uint32(0); i< me.DestUsr_Tl; i++ {
		me.Dest_Terminal_Id = append(me.Dest_Terminal_Id, r.ReadFixedString(32))
	}
	me.Msg_Length = r.ReadUint8()
	me.Msg_Content = r.ReadBytes(int(me.Msg_Length))
	me.Msg_Src = r.ReadFixedString(21)
	me.Src_Id = r.ReadFixedString(21)
	me.Service_Id = r.ReadFixedString(10)
	me.LinkID = r.ReadFixedString(20)
	me.Msg_Level = r.ReadUint8()
	me.Fee_UserType = r.ReadUint8()
	me.Fee_Terminal_Id = r.ReadFixedString(32)
	me.Fee_Terminal_Type = r.ReadUint8()
	me.TP_Pid = r.ReadUint8()
	me.TP_Udhi = r.ReadUint8()
	me.FeeType = r.ReadFixedString(2)
	me.FeeCode = r.ReadFixedString(6)
	me.Dest_Terminal_Type = r.ReadUint8()
}

func (me*EmppSubmit) Write(w *protocol.DataWriter) {
	w.WriteUint64(me.Msg_Id )
	w.WriteUint8(me.Pk_Total)
	w.WriteUint8(me.Pk_Number)
	w.WriteUint8(me.Registered_Delivery)
	w.WriteUint8(me.Msg_Fmt)
	w.WriteFixedString(me.Valid_Time, 17)
	w.WriteFixedString(me.At_Time, 17)
	me.DestUsr_Tl = uint32(len(me.Dest_Terminal_Id))
	w.WriteUint32(me.DestUsr_Tl)
	for i:= uint32(0); i< me.DestUsr_Tl; i++ {
		w.WriteFixedString( me.Dest_Terminal_Id[i], 32 )
	}
	me.Msg_Length = uint8(int(len(me.Msg_Content)))
	w.WriteUint8(me.Msg_Length)
	w.WriteBytes(me.Msg_Content)
	w.WriteFixedString(me.Msg_Src, 21)
	w.WriteFixedString(me.Src_Id, 21)
	w.WriteFixedString(me.Service_Id, 21)
	w.WriteFixedString(me.LinkID, 20)
	w.WriteUint8(me.Msg_Level)
	w.WriteUint8(me.Fee_UserType)
	w.WriteFixedString(me.Fee_Terminal_Id, 32)
	w.WriteUint8(me.Fee_Terminal_Type)
	w.WriteUint8(me.TP_Pid)
	w.WriteUint8(me.TP_Udhi)
	w.WriteFixedString(me.FeeType, 2)
	w.WriteFixedString(me.FeeCode, 6)
	w.WriteUint8(me.Dest_Terminal_Type)
}

func (me*EmppSubmit) AddDestTerminalId(mobile string) {
	me.Dest_Terminal_Id = append(me.Dest_Terminal_Id, mobile)
}
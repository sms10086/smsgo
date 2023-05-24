// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*EmppDeliver)(nil)
var _ protocol.PacketReader = (*EmppDeliver)(nil)

type EmppDeliver struct {
	Msg_Id []byte
	Dest_Id string
	Service_Id string
	TP_Pid uint8
	TP_Udhi uint8
	Msg_Fmt uint8
	Src_Terminal_Id string
	Src_Terminal_Type uint8
	Registered_Delivery uint8
	Msg_Length uint8
	Msg_Content []byte
	LinkID string
}

func (me*EmppDeliver) Read(r *protocol.DataReader) {
	me.Msg_Id = r.ReadBytes(10)
	me.Dest_Id = r.ReadFixedString(21)
	me.Service_Id = r.ReadFixedString(10)
	me.TP_Pid = r.ReadUint8()
	me.TP_Udhi = r.ReadUint8()
	me.Msg_Fmt = r.ReadUint8()
	me.Src_Terminal_Id = r.ReadFixedString(32)
	me.Src_Terminal_Type = r.ReadUint8()
	me.Registered_Delivery = r.ReadUint8()
	me.Msg_Length = r.ReadUint8()
	me.Msg_Content = r.ReadBytes(int(me.Msg_Length))
	me.LinkID = r.ReadFixedString(20)
}

func (me*EmppDeliver) Write(w *protocol.DataWriter) {
	w.WriteFixedBytes(me.Msg_Id, 10 )
	w.WriteFixedString(me.Dest_Id, 21)
	w.WriteFixedString(me.Service_Id, 10)
	w.WriteUint8(me.TP_Pid)
	w.WriteUint8(me.TP_Udhi)
	w.WriteUint8(me.Msg_Fmt)
	w.WriteFixedString(me.Src_Terminal_Id, 32)
	w.WriteUint8(me.Src_Terminal_Type)
	w.WriteUint8(me.Registered_Delivery)
	me.Msg_Length = uint8(len(me.Msg_Content))
	w.WriteUint8(me.Msg_Length)
	w.WriteBytes(me.Msg_Content)
	w.WriteFixedString(me.LinkID, 20)
}
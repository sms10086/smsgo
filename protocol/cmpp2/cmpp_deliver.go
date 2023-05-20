// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*CmppDeliver)(nil)
var _ protocol.PacketReader = (*CmppDeliver)(nil)

type CmppDeliver struct {
	Msg_Id uint64
	Dest_Id string
	Service_Id string
	TP_Pid uint8
	TP_Udhi uint8
	Msg_Fmt uint8
	Src_Terminal_Id string
	Registered_Delivery uint8
	Msg_Length uint8
	Msg_Content []byte
	Reserve []byte
}

func (me*CmppDeliver) Read(r *protocol.DataReader) {
	me.Msg_Id = r.ReadUint64()
	me.Dest_Id = r.ReadFixedString(21)
	me.Service_Id = r.ReadFixedString(10)
	me.TP_Pid = r.ReadUint8()
	me.TP_Udhi = r.ReadUint8()
	me.Msg_Fmt = r.ReadUint8()
	me.Src_Terminal_Id = r.ReadFixedString(21)
	me.Registered_Delivery = r.ReadUint8()
	me.Msg_Length = r.ReadUint8()
	me.Msg_Content = r.ReadBytes(int(me.Msg_Length))
	me.Reserve = r.ReadBytes(8)
}

func (me*CmppDeliver) Write(w *protocol.DataWriter) {
	w.WriteUint64(me.Msg_Id )
	w.WriteFixedString(me.Dest_Id, 21)
	w.WriteFixedString(me.Service_Id, 10)
	w.WriteUint8(me.TP_Pid)
	w.WriteUint8(me.TP_Udhi)
	w.WriteUint8(me.Msg_Fmt)
	w.WriteFixedString(me.Src_Terminal_Id, 21)
	w.WriteUint8(me.Registered_Delivery)
	me.Msg_Length = uint8(len(me.Msg_Content))
	w.WriteUint8(me.Msg_Length)
	w.WriteBytes(me.Msg_Content)
	w.WriteFixedBytes(me.Reserve, 8)
}
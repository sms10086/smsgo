// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindTransmitter)(nil)
var _ protocol.PacketReader = (*BindTransmitter)(nil)

type BindTransmitter struct {
	System_Id string
	Password string
	System_Type string
	Interface_Version uint8
	Addr_Ton uint8
	Addr_Npi uint8
	Address_Range string
}

func (me*BindTransmitter) Read(r *protocol.DataReader) {
	me.System_Id = r.ReadCString(16)
	me.Password = r.ReadCString(9)
	me.System_Type = r.ReadCString(13)
	me.Interface_Version = r.ReadUint8()
	me.Addr_Ton = r.ReadUint8()
	me.Addr_Npi = r.ReadUint8()
	me.Address_Range = r.ReadCString(41)
}

func (me*BindTransmitter) Write(w *protocol.DataWriter) {
	w.WriteCString(me.System_Id, 16)
	w.WriteCString(me.Password, 9)
	w.WriteCString(me.System_Type, 13)
	w.WriteUint8(me.Interface_Version)
	w.WriteUint8(me.Addr_Ton)
	w.WriteUint8(me.Addr_Npi)
	w.WriteCString(me.Address_Range, 41)
}

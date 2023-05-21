// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*DataSM)(nil)
var _ protocol.PacketReader = (*DataSM)(nil)

type DataSM struct {
	Service_Type string
	Source_Addr_Ton uint8
	Source_Addr_Npi uint8
	Source_Addr string
	Dest_Addr_Ton uint8
	Dest_Addr_Npi uint8
	Destination_Addr string
	Esm_Class uint8
	Registered_Delivery uint8
	Data_Coding uint8
}

func (me*DataSM) Read(r *protocol.DataReader) {
	me.Service_Type = r.ReadCString(6)
	me.Source_Addr_Ton = r.ReadUint8()
	me.Source_Addr_Npi = r.ReadUint8()
	me.Source_Addr = r.ReadCString(21)
	me.Dest_Addr_Ton = r.ReadUint8()
	me.Dest_Addr_Npi = r.ReadUint8()
	me.Destination_Addr = r.ReadCString(21)
	me.Esm_Class = r.ReadUint8()
	me.Registered_Delivery = r.ReadUint8()
	me.Data_Coding = r.ReadUint8()
}

func (me*DataSM) Write(w *protocol.DataWriter) {
	w.WriteCString(me.Service_Type, 6)
	w.WriteUint8(me.Source_Addr_Ton)
	w.WriteUint8(me.Source_Addr_Npi)
	w.WriteCString(me.Source_Addr, 21)
	w.WriteUint8(me.Dest_Addr_Ton)
	w.WriteUint8(me.Dest_Addr_Npi)
	w.WriteCString(me.Destination_Addr, 21)
	w.WriteUint8(me.Esm_Class)
	w.WriteUint8(me.Registered_Delivery)
	w.WriteUint8(me.Data_Coding)
}

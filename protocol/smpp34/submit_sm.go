// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*SubmitSM)(nil)
var _ protocol.PacketReader = (*SubmitSM)(nil)

type SubmitSM struct {
	Service_Type string
	Source_Addr_Ton uint8
	Source_Addr_Npi uint8
	Source_Addr string
	Dest_Addr_Ton uint8
	Dest_Addr_Npi uint8
	Destination_Addr string
	Esm_Class uint8
	Protocol_Id uint8
	Priority_Flag uint8
	Schedule_Delivery_Time string
	Validity_Period string
	Registered_Delivery uint8
	Replace_If_Present_Flag uint8
	Data_Coding uint8
	Sm_Default_Msg_Id uint8
	Sm_Length uint8
	Short_Message []byte
}

func (me*SubmitSM) Read(r *protocol.DataReader) {
	me.Service_Type = r.ReadCString(6)
	me.Source_Addr_Ton = r.ReadUint8()
	me.Source_Addr_Npi = r.ReadUint8()
	me.Source_Addr = r.ReadCString(21)
	me.Dest_Addr_Ton = r.ReadUint8()
	me.Dest_Addr_Npi = r.ReadUint8()
	me.Destination_Addr = r.ReadCString(21)
	me.Esm_Class = r.ReadUint8()
	me.Protocol_Id = r.ReadUint8()
	me.Priority_Flag = r.ReadUint8()
	me.Schedule_Delivery_Time = r.ReadCString(17)
	me.Validity_Period = r.ReadCString(17)
	me.Registered_Delivery = r.ReadUint8()
	me.Replace_If_Present_Flag = r.ReadUint8()
	me.Data_Coding = r.ReadUint8()
	me.Sm_Default_Msg_Id = r.ReadUint8()
	me.Sm_Length = r.ReadUint8()
	me.Short_Message = r.ReadBytes(int(me.Sm_Length))
}

func (me*SubmitSM) Write(w *protocol.DataWriter) {
	w.WriteCString(me.Service_Type, 6)
	w.WriteUint8(me.Source_Addr_Ton)
	w.WriteUint8(me.Source_Addr_Npi)
	w.WriteCString(me.Source_Addr, 21)
	w.WriteUint8(me.Dest_Addr_Ton)
	w.WriteUint8(me.Dest_Addr_Npi)
	w.WriteCString(me.Destination_Addr, 21)
	w.WriteUint8(me.Esm_Class)
	w.WriteUint8(me.Protocol_Id)
	w.WriteUint8(me.Priority_Flag)
	w.WriteCString(me.Schedule_Delivery_Time, 17)
	w.WriteCString(me.Validity_Period, 17)
	w.WriteUint8(me.Registered_Delivery)
	w.WriteUint8(me.Replace_If_Present_Flag)
	w.WriteUint8(me.Data_Coding)
	w.WriteUint8(me.Sm_Default_Msg_Id)
	me.Sm_Length = uint8(len(me.Short_Message))
	w.WriteUint8(me.Sm_Length)
	w.WriteBytes(me.Short_Message)
}

// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Charge)(nil)
var _ protocol.PacketReader = (*Charge)(nil)

type Charge struct {
	Charge_Ratio uint32
	Charge_Type uint32
	Charge_Resouce string
	Charge_Location uint8
}

func (me*Charge) Read(r *protocol.DataReader) {
	me.Charge_Ratio = r.ReadUint32LE()
	me.Charge_Type = r.ReadUint32LE()
	me.Charge_Resouce = r.ReadFixedString(21)
	me.Charge_Location = r.ReadUint8()
}

func (me*Charge) Write(w *protocol.DataWriter) {
	w.WriteUint32LE(me.Charge_Ratio)
	w.WriteUint32LE(me.Charge_Type)
	w.WriteFixedCString(me.Charge_Resouce, 21)
	w.WriteUint8(me.Charge_Location)
}

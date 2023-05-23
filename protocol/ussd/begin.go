// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Begin)(nil)
var _ protocol.PacketReader = (*Begin)(nil)

type Begin struct {
	Ussd_Version uint8
	Ussd_Op_Type uint8
	MsIsdn string
	Service_Code string
	Code_Scheme uint8
	Ussd_Content []byte
}

func (me*Begin) Read(r *protocol.DataReader) {
	me.Ussd_Version = r.ReadUint8()
	me.Ussd_Op_Type = r.ReadUint8()
	me.MsIsdn = r.ReadFixedString(21)
	me.Service_Code = r.ReadFixedString(21)
	me.Code_Scheme = r.ReadUint8()
}

func (me*Begin) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Ussd_Version)
	w.WriteUint8(me.Ussd_Op_Type)
	w.WriteFixedCString(me.MsIsdn, 21)
	w.WriteFixedCString(me.Service_Code, 21)
	w.WriteUint8(me.Code_Scheme)
	w.WriteBytes(me.Ussd_Content)
}

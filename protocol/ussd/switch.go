// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Switch)(nil)
var _ protocol.PacketReader = (*Switch)(nil)

type Switch struct {
	Switch_Mode uint8
	MsIsdn string
	Org_Service_Code string
	Dest_Service_Code string
	Code_Scheme uint8
	Ussd_Content []byte
}

func (me*Switch) Read(r *protocol.DataReader) {
	me.Switch_Mode = r.ReadUint8()
	me.MsIsdn = r.ReadFixedString(21)
	me.Org_Service_Code = r.ReadFixedString(21)
	me.Dest_Service_Code = r.ReadFixedString(21)
	me.Code_Scheme = r.ReadUint8()
}

func (me*Switch) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Switch_Mode)
	w.WriteFixedCString(me.MsIsdn, 21)
	w.WriteFixedCString(me.Org_Service_Code, 21)
	w.WriteFixedCString(me.Dest_Service_Code, 21)
	w.WriteUint8(me.Code_Scheme)
	w.WriteBytes(me.Ussd_Content)
}

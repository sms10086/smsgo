// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Bind)(nil)
var _ protocol.PacketReader = (*Bind)(nil)

type Bind struct {
	System_ID string
	Password string
	System_Type string
	Interface_Version uint8
}

func (me*Bind) Read(r *protocol.DataReader) {
	me.System_ID = r.ReadFixedString(11)
	me.Password = r.ReadFixedString(9)
	me.System_Type = r.ReadFixedString(13)
	me.Interface_Version = r.ReadUint8()
}

func (me*Bind) Write(w *protocol.DataWriter) {
	w.WriteFixedCString(me.System_ID, 11)
	w.WriteFixedCString(me.Password, 9)
	w.WriteFixedCString(me.System_Type, 13)
	w.WriteUint8(me.Interface_Version)
}

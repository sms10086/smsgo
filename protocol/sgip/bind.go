// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*Bind)(nil)
var _ protocol.PacketReader = (*Bind)(nil)

type Bind struct {
	Login_Type uint8
	Login_Name string
	Login_Password string
	Reserve string
}

func (me*Bind) Read(r *protocol.DataReader) {
	me.Login_Type = r.ReadUint8()
	me.Login_Name = r.ReadFixedString(16)
	me.Login_Password = r.ReadFixedString(16)
	me.Reserve = r.ReadFixedString(8)
}

func (me*Bind) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Login_Type)
	w.WriteFixedString(me.Login_Name, 16)
	w.WriteFixedString(me.Login_Password, 16)
	w.WriteFixedString(me.Reserve, 8)
}

func NewBind(username, password string) *Bind {
	c := new(Bind)
	c.Login_Type = 1
	c.Login_Name = username
	c.Login_Password = password
	return c
}
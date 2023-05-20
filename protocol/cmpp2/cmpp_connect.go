// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*CmppConnect)(nil)
var _ protocol.PacketReader = (*CmppConnect)(nil)

type CmppConnect struct {
	Source_Addr string
	AuthenticatorSource []byte
	Version uint8
	Timestamp uint32
}

func (me*CmppConnect) Read(r *protocol.DataReader) {
	me.Source_Addr = r.ReadFixedString(6)
	me.AuthenticatorSource = r.ReadBytes(16)
	me.Version = r.ReadUint8()
	me.Timestamp = r.ReadUint32()
}

func (me*CmppConnect) Write(w *protocol.DataWriter) {
	w.WriteFixedString(me.Source_Addr, 6)
	w.WriteBytes(me.AuthenticatorSource)
	w.WriteUint8(me.Version)
	w.WriteUint32(me.Timestamp)
}
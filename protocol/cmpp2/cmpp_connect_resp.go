// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*CmppConnectResp)(nil)
var _ protocol.PacketReader = (*CmppConnectResp)(nil)

type CmppConnectResp struct {
	Status uint8
	AuthenticatorISMG []byte
	Version uint8
}

func (me*CmppConnectResp) Read(r *protocol.DataReader) {
	me.Status = r.ReadUint8()
	me.AuthenticatorISMG = r.ReadBytes(16)
	me.Version = r.ReadUint8()
}

func (me*CmppConnectResp) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Status)
	w.WriteBytes(me.AuthenticatorISMG)
	w.WriteUint8(me.Version)
}
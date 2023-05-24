// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*EmppConnectResp)(nil)
var _ protocol.PacketReader = (*EmppConnectResp)(nil)

type EmppConnectResp struct {
	Status uint32
	AuthenticatorESMP []byte
	Version uint8
	Ability uint32
}

func (me*EmppConnectResp) Read(r *protocol.DataReader) {
	me.Status = r.ReadUint32()
	me.AuthenticatorESMP = r.ReadBytes(16)
	me.Version = r.ReadUint8()
	me.Ability = r.ReadUint32()
}

func (me*EmppConnectResp) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.Status)
	w.WriteBytes(me.AuthenticatorESMP)
	w.WriteUint8(me.Version)
	w.WriteUint32(me.Ability)
}
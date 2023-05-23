// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smgp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*LoginResp)(nil)
var _ protocol.PacketReader = (*LoginResp)(nil)

type LoginResp struct {
	Status uint32
	AuthenticatorServer []byte
	ServerVersion uint8
}

func (me*LoginResp) Read(r *protocol.DataReader) {
	me.Status = r.ReadUint32()
	me.AuthenticatorServer = r.ReadBytes(16)
	me.ServerVersion = r.ReadUint8()
}

func (me*LoginResp) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.Status)
	w.WriteBytes(me.AuthenticatorServer)
	w.WriteUint8(me.ServerVersion)
}
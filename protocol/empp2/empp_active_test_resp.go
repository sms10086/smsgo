// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*EmppActiveTestResp)(nil)
var _ protocol.PacketReader = (*EmppActiveTestResp)(nil)

type EmppActiveTestResp struct {
	Reserved uint8
}

func (me*EmppActiveTestResp) Read(r *protocol.DataReader) {
	me.Reserved = r.ReadUint8()
}

func (me*EmppActiveTestResp) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Reserved)
}
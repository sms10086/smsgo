// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*CmppActiveTestResp)(nil)
var _ protocol.PacketReader = (*CmppActiveTestResp)(nil)

type CmppActiveTestResp struct {
	Reserved uint8
}

func (me*CmppActiveTestResp) Read(r *protocol.DataReader) {
	me.Reserved = r.ReadUint8()
}

func (me*CmppActiveTestResp) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Reserved)
}
// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindTransmitterResp)(nil)
var _ protocol.PacketReader = (*BindTransmitterResp)(nil)

type BindTransmitterResp struct {
	System_Id string
}

func (me*BindTransmitterResp) Read(r *protocol.DataReader) {
	me.System_Id = r.ReadCString(16)
}

func (me*BindTransmitterResp) Write(w *protocol.DataWriter) {
	w.WriteCString(me.System_Id, 16)
}

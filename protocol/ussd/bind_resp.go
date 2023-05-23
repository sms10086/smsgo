// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindResp)(nil)
var _ protocol.PacketReader = (*BindResp)(nil)

type BindResp struct {
	System_ID string
}

func (me*BindResp) Read(r *protocol.DataReader) {
	me.System_ID = r.ReadFixedString(11)
}

func (me*BindResp) Write(w *protocol.DataWriter) {
	w.WriteFixedCString(me.System_ID, 11)
}
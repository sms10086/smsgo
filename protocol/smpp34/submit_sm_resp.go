// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*SubmitSMResp)(nil)
var _ protocol.PacketReader = (*SubmitSMResp)(nil)

type SubmitSMResp struct {
	Message_Id string
}

func (me*SubmitSMResp) Read(r *protocol.DataReader) {
	me.Message_Id = r.ReadCString(65)
}

func (me*SubmitSMResp) Write(w *protocol.DataWriter) {
	w.WriteCString(me.Message_Id, 65)
}

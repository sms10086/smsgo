// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cngp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*SubmitResp)(nil)
var _ protocol.PacketReader = (*SubmitResp)(nil)

type SubmitResp struct {
	MsgID []byte
	TLVs protocol.TLVs
}

func (me*SubmitResp) Read(r *protocol.DataReader) {
	me.MsgID = r.ReadBytes(10)
}

func (me*SubmitResp) Write(w *protocol.DataWriter) {
	w.WriteFixedBytes(me.MsgID, 10)
	w.WriteTLVs(me.TLVs)
}

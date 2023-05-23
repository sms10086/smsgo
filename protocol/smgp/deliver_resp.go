// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smgp

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*DeliverResp)(nil)
var _ protocol.PacketReader = (*DeliverResp)(nil)

type DeliverResp struct {
	MsgID []byte
	Status uint32
}

func (me*DeliverResp) Read(r *protocol.DataReader) {
	me.MsgID = r.ReadBytes(10)
	me.Status = r.ReadUint32()
}

func (me*DeliverResp) Write(w *protocol.DataWriter) {
	w.WriteFixedBytes(me.MsgID, 10)
	w.WriteUint32(me.Status)
}

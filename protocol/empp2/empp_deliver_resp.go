// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*EmppDeliverResp)(nil)
var _ protocol.PacketReader = (*EmppDeliverResp)(nil)

type EmppDeliverResp struct {
	Msg_Id []byte
	Result uint32
}

func (me*EmppDeliverResp) Read(r *protocol.DataReader) {
	me.Msg_Id = r.ReadBytes(10)
	me.Result = r.ReadUint32()
}

func (me*EmppDeliverResp) Write(w *protocol.DataWriter) {
	w.WriteFixedBytes(me.Msg_Id, 10)
	w.WriteUint32(me.Result)
}
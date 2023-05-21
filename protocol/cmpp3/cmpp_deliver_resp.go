// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp3

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*CmppDeliverResp)(nil)
var _ protocol.PacketReader = (*CmppDeliverResp)(nil)

type CmppDeliverResp struct {
	Msg_Id uint64
	Result uint32
}

func (me*CmppDeliverResp) Read(r *protocol.DataReader) {
	me.Msg_Id = r.ReadUint64()
	me.Result = r.ReadUint32()
}

func (me*CmppDeliverResp) Write(w *protocol.DataWriter) {
	w.WriteUint64(me.Msg_Id)
	w.WriteUint32(me.Result)
}
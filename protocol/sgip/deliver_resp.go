// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*DeliverResp)(nil)
var _ protocol.PacketReader = (*DeliverResp)(nil)

type DeliverResp struct {
	Result uint8
	Reserve string
}

func (me*DeliverResp) Read(r *protocol.DataReader) {
	me.Result = r.ReadUint8()
	me.Reserve = r.ReadFixedString(8)
}

func (me*DeliverResp) Write(w *protocol.DataWriter) {
	w.WriteUint8(me.Result)
	w.WriteFixedString(me.Reserve, 8)
}

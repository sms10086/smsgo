// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*DeliverSMResp)(nil)
var _ protocol.PacketReader = (*DeliverSMResp)(nil)

type DeliverSMResp struct {
}

func (me*DeliverSMResp) Read(r *protocol.DataReader) {
}

func (me*DeliverSMResp) Write(w *protocol.DataWriter) {
	w.WriteUint8(0)
}

// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindReceiverResp)(nil)
var _ protocol.PacketReader = (*BindReceiverResp)(nil)

type BindReceiverResp BindTransmitterResp

func (me*BindReceiverResp) Read(r *protocol.DataReader) {
	(*BindTransmitterResp)(me).Read(r)
}

func (me*BindReceiverResp) Write(w *protocol.DataWriter) {
	(*BindTransmitterResp)(me).Write(w)
}

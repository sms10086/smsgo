// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindTransceiverResp)(nil)
var _ protocol.PacketReader = (*BindTransceiverResp)(nil)

type BindTransceiverResp BindTransmitterResp

func (me*BindTransceiverResp) Read(r *protocol.DataReader) {
	(*BindTransceiverResp)(me).Read(r)
}

func (me*BindTransceiverResp) Write(w *protocol.DataWriter) {
	(*BindTransceiverResp)(me).Write(w)
}

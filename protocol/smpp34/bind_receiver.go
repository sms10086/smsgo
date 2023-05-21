// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindReceiver)(nil)
var _ protocol.PacketReader = (*BindReceiver)(nil)

type BindReceiver BindTransmitter

func (me*BindReceiver) Read(r *protocol.DataReader) {
	(*BindTransmitter)(me).Read(r)
}

func (me*BindReceiver) Write(w *protocol.DataWriter) {
	(*BindTransmitter)(me).Write(w)
}
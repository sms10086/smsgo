// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*BindTransceiver)(nil)
var _ protocol.PacketReader = (*BindTransceiver)(nil)

type BindTransceiver BindTransmitter

func (me*BindTransceiver) Read(r *protocol.DataReader) {
	(*BindTransceiver)(me).Read(r)
}

func (me*BindTransceiver) Write(w *protocol.DataWriter) {
	(*BindTransceiver)(me).Write(w)
}
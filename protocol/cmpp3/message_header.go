// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp3

import (
	"smsgo/protocol"
	"smsgo/protocol/cmpp2"
)

var _ protocol.PacketWriter = (*MessageHeader)(nil)
var _ protocol.PacketReader = (*MessageHeader)(nil)

type MessageHeader cmpp2.MessageHeader

func (me*MessageHeader) Read(r *protocol.DataReader) {
	(*cmpp2.MessageHeader)(me).Read(r)
}

func (me*MessageHeader) Write(w *protocol.DataWriter) {
	(*cmpp2.MessageHeader)(me).Write(w)
}
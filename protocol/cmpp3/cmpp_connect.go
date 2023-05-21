// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp3

import (
	"smsgo/protocol"
	"smsgo/protocol/cmpp2"
)

var _ protocol.PacketWriter = (*CmppConnect)(nil)
var _ protocol.PacketReader = (*CmppConnect)(nil)


type CmppConnect cmpp2.CmppConnect

func (me*CmppConnect) Read(r *protocol.DataReader) {
	(*cmpp2.CmppConnect)(me).Read(r)
}

func (me*CmppConnect) Write(w *protocol.DataWriter) {
	(*cmpp2.CmppConnect)(me).Write(w)
}
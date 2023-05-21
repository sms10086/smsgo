// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp3

import (
	"smsgo/protocol"
	"smsgo/protocol/cmpp2"
)

var _ protocol.PacketWriter = (*CmppActiveTestResp)(nil)
var _ protocol.PacketReader = (*CmppActiveTestResp)(nil)

type CmppActiveTestResp cmpp2.CmppActiveTestResp

func (me*CmppActiveTestResp) Read(r *protocol.DataReader) {
	(*cmpp2.CmppActiveTestResp)(me).Read(r)
}

func (me*CmppActiveTestResp) Write(w *protocol.DataWriter) {
	(*cmpp2.CmppActiveTestResp)(me).Write(w)
}
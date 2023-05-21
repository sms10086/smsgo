// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smpp34

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*DeliverSM)(nil)
var _ protocol.PacketReader = (*DeliverSM)(nil)

type DeliverSM SubmitSM

func (me*DeliverSM) Read(r *protocol.DataReader) {
	(*SubmitSM)(me).Read(r)
}

func (me*DeliverSM) Write(w *protocol.DataWriter) {
	(*SubmitSM)(me).Write(w)
}
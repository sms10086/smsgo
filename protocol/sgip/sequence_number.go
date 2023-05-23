// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package sgip

import (
	"smsgo/protocol"
)

var _ protocol.PacketWriter = (*SequenceNumber)(nil)
var _ protocol.PacketReader = (*SequenceNumber)(nil)

type SequenceNumber struct {
	Node uint32
	Time uint32
	Sequence uint32
}

func (me*SequenceNumber) Read(r *protocol.DataReader) {
	me.Node = r.ReadUint32()
	me.Time = r.ReadUint32()
	me.Sequence = r.ReadUint32()
}

func (me*SequenceNumber) Write(w *protocol.DataWriter) {
	w.WriteUint32(me.Node)
	w.WriteUint32(me.Time)
	w.WriteUint32(me.Sequence)
}

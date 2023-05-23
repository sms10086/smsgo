// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strings"
)

var _ PacketWriter = (*TLV)(nil)
var _ PacketReader = (*TLV)(nil)

type TLV struct {
	Tag uint16
	Length uint16
	Value []byte
}

func NewTLV(tag int) *TLV {
	var t = new(TLV)
	t.Tag = uint16(tag)
	return t
}

func (me*TLV) Read(r *DataReader) {
	me.Tag = r.ReadUint16()
	me.Length = r.ReadUint16()
	me.Value = r.ReadBytes(int(me.Length))
}

func (me*TLV) Write(w *DataWriter) {
	w.WriteUint16(me.Tag)
	me.Length = uint16(len(me.Value))
	w.WriteUint16(me.Length)
	w.WriteBytes(me.Value)
}

func (me*TLV) GetByte() int {
	if len(me.Value) == 0 {
		return 0
	}
	return int(me.Value[0])
}

func (me*TLV) SetByte(val int) {
	me.Length = 1
	me.Value = make([]byte, 1)
	me.Value[0] = uint8(val)
}

func (me*TLV) GetString() string {
	return strings.TrimSpace(string(me.Value))
}

func (me*TLV) SetString(s string) {
	me.Value = []byte(s)
	me.Length = uint16(len(me.Value))
}

type TLVs map[int]*TLV

func NewTLVs() TLVs {
	return make(map[int]*TLV)
}

func (me TLVs) PutByte(tag int, val int) {
	var t = NewTLV(tag)
	t.SetByte(val)
	me[tag] = t
}

func (me TLVs) GetByte(tag int) int {
	if t := me[tag]; t != nil {
		return t.GetByte()
	}
	return 0
}

func (me TLVs) PutString(tag int, s string) {
	var t = NewTLV(tag)
	t.SetString(s)
	me[tag] = t
}

func (me TLVs) GetString(tag int) string {
	if t := me[tag]; t != nil {
		return t.GetString()
	}
	return ""
}
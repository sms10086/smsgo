// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

import (
	"smsgo/protocol"
	"time"
	"crypto/md5"
	"strconv"
)

var _ protocol.PacketWriter = (*EmppConnect)(nil)
var _ protocol.PacketReader = (*EmppConnect)(nil)

type EmppConnect struct {
	AccountId string
	AuthenticatorSource []byte
	Version uint8
	Timestamp uint32
}

func (me*EmppConnect) Read(r *protocol.DataReader) {
	me.AccountId = r.ReadFixedString(21)
	me.AuthenticatorSource = r.ReadBytes(16)
	me.Version = r.ReadUint8()
	me.Timestamp = r.ReadUint32()
}

func (me*EmppConnect) Write(w *protocol.DataWriter) {
	w.WriteFixedString(me.AccountId, 21)
	w.WriteBytes(me.AuthenticatorSource)
	w.WriteUint8(me.Version)
	w.WriteUint32(me.Timestamp)
}

func (me*EmppConnect) SetAuthentication(username, password string) {
	
	ts := time.Now().Format("20060102150405")[4:]
	i, _ := strconv.ParseInt(ts, 10, 32)
	md := md5.New()
	md.Write([]byte(username))
	md.Write(make([]byte, 9))
	md.Write([]byte(password))
	md.Write([]byte(ts))

	me.AccountId = username
	me.AuthenticatorSource = md.Sum(nil)
	me.Timestamp = uint32(i)
}

func NewEmppConnect(username, password string, ver int) *EmppConnect {
	c := new(EmppConnect)
	c.Version = uint8(ver)
	c.SetAuthentication(username, password)
	return c
}
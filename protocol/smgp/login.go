// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package smgp

import (
	"smsgo/protocol"
	"time"
	"crypto/md5"
	"strconv"
)

var _ protocol.PacketWriter = (*Login)(nil)
var _ protocol.PacketReader = (*Login)(nil)

type Login struct {
	ClientID string
	AuthenticatorClient []byte
	LoginMode uint8
	TimeStamp uint32
	ClientVersion uint8
}

func (me*Login) Read(r *protocol.DataReader) {
	me.ClientID = r.ReadFixedString(8)
	me.AuthenticatorClient = r.ReadBytes(16)
	me.LoginMode = r.ReadUint8()
	me.TimeStamp = r.ReadUint32()
	me.ClientVersion = r.ReadUint8()
}

func (me*Login) Write(w *protocol.DataWriter) {
	w.WriteFixedString(me.ClientID, 6)
	w.WriteBytes(me.AuthenticatorClient)
	w.WriteUint8(me.LoginMode)
	w.WriteUint32(me.TimeStamp)
	w.WriteUint8(me.ClientVersion)
}

func (me*Login) SetAuthentication(username, password string) {
	
	ts := time.Now().Format("20060102150405")[4:]
	i, _ := strconv.ParseInt(ts, 10, 32)
	md := md5.New()
	md.Write([]byte(username))
	md.Write(make([]byte, 7))
	md.Write([]byte(password))
	md.Write([]byte(ts))
	
	me.ClientID = username
	me.AuthenticatorClient = md.Sum(nil)
	me.TimeStamp = uint32(i)
}

func NewLogin(username, password string, ver int) *Login {
	c := new(Login)
	c.ClientVersion = uint8(ver)
	c.LoginMode = LOGINMODE_TRANSMIT_MODE
	c.SetAuthentication(username, password)
	return c
}
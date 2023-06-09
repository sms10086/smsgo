// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package protocol

import (
	"io"
	"bytes"
	"encoding/binary"
)

type DataWriter struct {
	w io.Writer
	buf bytes.Buffer
}

func (me*DataWriter) WriteUint8(i uint8) {
	me.buf.WriteByte(i)
}

func (me*DataWriter) WriteUint16(i uint16) {
	me.buf.WriteByte( uint8(i>>8) )
	me.buf.WriteByte( uint8(i) )
}

func (me*DataWriter) WriteFixedString(s string, n int) {
	d := []byte(s)
	me.WriteFixedBytes(d, n)
}

func (me*DataWriter) WriteFixedCString(s string, n int) {
	d := []byte(s)
	if len(d) > n-1 {
		d = d[0:n-1]
	}
	me.WriteFixedBytes(d, n)
}

func (me*DataWriter) WriteUint32(v uint32) {
	d := make([]byte, 4)
	binary.BigEndian.PutUint32(d, v)
	me.buf.Write(d)
}

func (me*DataWriter) WriteUint32LE(v uint32) {
	d := make([]byte, 4)
	binary.LittleEndian.PutUint32(d, v)
	me.buf.Write(d)
}

func (me*DataWriter) WriteUint64(v uint64) {
	d := make([]byte, 8)
	binary.BigEndian.PutUint64(d, v)
	me.buf.Write(d)
}

func (me*DataWriter) WriteBytes(buf []byte) {
	if len(buf) == 0 {
		return
	}
	me.buf.Write(buf)
}

func (me*DataWriter) WriteFixedBytes(buf[]byte, n int) {
	i := len(buf)
	if i >= n {
		me.buf.Write(buf[0:n])
	} else {
		if i > 0 {
			me.buf.Write(buf)
		}
		for ; i< n; i++ {
			me.buf.WriteByte(0)
		}
	}
}

func (me*DataWriter) WriteCString(s string, n int) {
	buf := []byte(s)
	if len(buf) >= n {
		buf = buf[0:n-1]
	}
	me.buf.Write(buf)
	me.buf.WriteByte(0)
}

func (me*DataWriter) WriteTLVs(tlvs TLVs) {
	for _, t := range tlvs {
		t.Write(me)
	}
}

func (me*DataWriter) Write(p PacketWriter) error {
	p.Write(me)
	_, err := me.Send()
	return err
}

func (me*DataWriter) Send() (int, error) {
	d := me.buf.Bytes()
	binary.BigEndian.PutUint32(d, uint32(len(d)))
	n, err := me.w.Write(d)
	me.buf.Reset()
	return n, err
}

func NewDataWriter(writer io.Writer) *DataWriter {
	w := new(DataWriter)
	w.w = writer
	return w
}
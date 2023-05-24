// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package protocol

import (
	"bytes"
	"io"
	"encoding/binary"
	"strings"
)

type DataReader struct {
	r io.Reader
	off int
	err error
}

func (me*DataReader) Reset() {
	me.off = 0
	me.err = nil
}

func (me*DataReader) ReadUint8() uint8 {
	if me.err != nil {
		return 0
	}
	buf := make([]byte, 1)
	if _, me.err = me.r.Read(buf); me.err != nil {
		return 0
	}
	me.off ++
	return buf[0]
}

func (me*DataReader) ReadUint16() uint16 {
	if me.err != nil {
		return 0
	}
	buf := make([]byte, 2)
	if _, me.err = me.r.Read(buf); me.err != nil {
		return 0
	}
	me.off = me.off + 2
	return binary.BigEndian.Uint16(buf)
}

func (me*DataReader) ReadFixedString(n int) string {
	if me.err != nil || n <= 0 {
		return ""
	}
	var buf = make([]byte, n)
	_, me.err = me.r.Read(buf)
	if me.err != nil {
		return ""
	}
	me.off = me.off + n
	return strings.TrimSpace(string(buf))
}

func (me*DataReader) ReadBytes(n int) []byte {
	if me.err != nil || n <= 0 {
		return nil
	}
	var buf = make([]byte, n)
	_, me.err = me.r.Read(buf)
	if me.err != nil {
		return nil
	}
	me.off = me.off + n
	return buf
}

func (me*DataReader) ReadUint32() uint32 {
	if me.err != nil {
		return 0
	}
	var buf = make([]byte, 4)
	_, me.err = me.r.Read(buf)
	if me.err != nil {
		return 0
	}
	me.off = me.off + 4
	return binary.BigEndian.Uint32(buf)
}

func (me*DataReader) ReadUint32LE() uint32 {
	if me.err != nil {
		return 0
	}
	var buf = make([]byte, 4)
	_, me.err = me.r.Read(buf)
	if me.err != nil {
		return 0
	}
	me.off = me.off + 4
	return binary.LittleEndian.Uint32(buf)
}

func (me*DataReader) ReadUint64() uint64 {
	if me.err != nil {
		return 0
	}
	var buf = make([]byte, 8)
	_, me.err = me.r.Read(buf)
	if me.err != nil {
		return 0
	}
	me.off = me.off + 8
	return binary.BigEndian.Uint64(buf)
}

func (me*DataReader) ReadCString(n int) string {
	if me.err != nil {
		return ""
	}
	var b bytes.Buffer
	for i:=0; i<n; i++ {
		v := me.ReadUint8()
		if v == 0 {
			break
		}
		b.WriteByte(v)
	}
	me.off = me.off + b.Len() + 1
	return b.String()
}

func (me*DataReader) ReadTLVs(packetSize int) (t TLVs) {
	t = NewTLVs()
	if me.err != nil {
		return
	}
	for me.off < packetSize {
		x := new(TLV)
		x.Read(me)
		if me.err != nil {
			break
		}
		t[int(x.Tag)] = x
	}
	return
}

func (me*DataReader) Error() error {
	return me.err
}

func (me*DataReader) Read(p PacketReader) error {
	p.Read(me)
	return me.err
}

func NewDataReader(r io.Reader) *DataReader {
	reader := new(DataReader)
	reader.r = r
	return reader
}
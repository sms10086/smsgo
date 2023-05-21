// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

import (
	"smsgo/protocol"
	"errors"
)

func WriteHeader(w *protocol.DataWriter, commandId uint32, seqId uint32) {
	h := new(MessageHeader)
	h.Command_Id = commandId
	h.Sequence_Id = seqId
	h.Write(w)
}

func WritePacket(w *protocol.DataWriter, commandId uint32, seqId uint32, body protocol.PacketWriter) error {
	WriteHeader(w, commandId, seqId)
	if body != nil {
		body.Write(w)
	}
	_, err := w.Send()
	return err
}

func ReadHeader(r *protocol.DataReader)(commandId uint32, seqId uint32, err error) {
	h := new(MessageHeader)
	err = r.Read(h)
	commandId = h.Command_Id
	seqId = h.Sequence_Id
	return
}

func ReadConnectResp(r *protocol.DataReader) (int, error) {
	cmd, _, _ := ReadHeader(r)
	if cmd != CMPP_CONNECT_RESP {
		return -1, errors.New("not CMPP_CONNECT_RESP")
	}
	resp := new(CmppConnectResp)
	resp.Read(r)
	return int(resp.Status), nil
}

func SendMessage(w *protocol.DataWriter, seqId uint32, shortCode, mobile string, msgFmt uint8, content[]byte) error {
	submit := new(CmppSubmit)
	submit.Registered_Delivery = 1
	submit.Msg_Fmt = msgFmt
	submit.Src_Id = shortCode
	submit.AddDestTerminalId(mobile)
	submit.Msg_Content = content
	return WritePacket(w, CMPP_SUBMIT, seqId, submit)
}

func ReadSubmitResp(r *protocol.DataReader) (result int, msgId uint64, err error) {
	submit_resp := new(CmppSubmitResp)
	submit_resp.Read(r)
	err = r.Error()
	result = int(submit_resp.Result)
	msgId = submit_resp.Msg_Id
	return
}
// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cngp

const (
	// SP->SMGW	登录请求
	LOGIN = 0x00000001
	// SP->SMGW	登录请求的应答
	LOGIN_RESP = 0X80000001
	// SP->SMGW	SP发送短消息请求
	SUBMIT = 0x00000002
	// SP->SMGW	SP发送短消息请求的应答
	SUBMIT_RESP = 0x80000002
	// SMGW->SP	SMGW发送短消息请求
	DELIVER = 0x00000003
	// SMGW->SP	SMGW发送短消息的应答
	DELIVER_RESP = 0x80000003
	// SP->SMGW	测试通信链路是否正常请求（由客户端发起，SP和SMGW可以通过定时发送此请求来维持连接）
	ACTIVE_TEST = 0x00000004
	// SP->SMGW	测试通信链路是否正常的应答
	ACTIVE_TEST_RESP = 0x80000004
	// SP->SMGW	退出请求
	EXIT = 0x00000006
	// SP->SMGW	退出请求的应答
	EXIT_RESP = 0x80000006
)
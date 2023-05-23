// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package ussd

const (
	// 请求连接
	BIND = 0x00000065
	// 请求连接应答
	BIND_RESP = 0x00000067
	// 终止连接
	UNBIND = 0x00000066
	// 终止连接应答
	UNBIND_RESP = 0x00000068
	// 握手请求
	ENQUIRE_LINK = 0x00000083
	// 握手应答
	ENQUIRE_LINK_RESP = 0x00000084
	// 开始USSD会话
	BEGIN = 0x0000006F
	// 继续USSD会话
	CONTINUE = 0x00000070
	// 结束USSD会话
	END = 0x00000071
	// 中止USSD会话 
	ABORT = 0x00000072
	// USSD会话转移
	SWITCH = 0x00000074
	// 计费指示
	CHARGEIND = 0x00000075
	// 计费指示应答
	CHARGEIND_RESP = 0x00000076
	// 开始转移USSD会话
	SWITCH_BEGIN = 0x00000077
)
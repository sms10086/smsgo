// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package empp2

const (
	// 请求连接
	EMPP_CONNECT = 0x00000001 
	// 请求连接应答
	EMPP_CONNECT_RESP = 0x80000001
	// 终止连接
	EMPP_TERMINATE = 0x00000002
	// 终止连接应答
	EMPP_TERMINATE_RESP = 0x80000002
	// 提交短信
	EMPP_SUBMIT = 0x00000004
	// 提交短信应答
	EMPP_SUBMIT_RESP = 0x80000004
	// 短信回复
	EMPP_DELIVER = 0x00000005
	// 回复短信应答
	EMPP_DELIVER_RESP = 0x80000005
	// 激活测试
	EMPP_ACTIVE_TEST = 0x00000008
	// 激活测试应答
	EMPP_ACTIVE_TEST_RESP = 0x80000008
)
// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

// 7.7.1 Command_Id 定义
const (
	// 请求连接
	CMPP_CONNECT        = 0x00000001 
	// 请求连接应答
	CMPP_CONNECT_RESP   = 0x80000001
	// 终止连接
	CMPP_TERMINATE      = 0x00000002
	// 终止连接应答
	CMPP_TERMINATE_RESP = 0x80000002
	// 提交短信
	CMPP_SUBMIT         = 0x00000004
	// 提交短信应答
	CMPP_SUBMIT_RESP    = 0x80000004
	// 短信下发
	CMPP_DELIVER        = 0x00000005
	// 下发短信应答
	CMPP_DELIVER_RESP   = 0x80000005
	// 发送短信状态查询
	CMPP_QUERY          = 0x00000006
	// 发送短信状态查询应答
	CMPP_QUERY_RESP     = 0x80000006
	// 删除短信
	CMPP_CANCEL         = 0x00000007
	// 删除短信应答
	CMPP_CANCEL_RESP    = 0x80000007
	// 激活测试
	CMPP_ACTIVE_TEST    = 0x00000008
	// 激活测试应答
	CMPP_ACTIVE_TEST_RESP = 0x80000008
	// 消息前转
	CMPP_FWD            = 0x00000009
	// 消息前转应答
	CMPP_FWD_RESP       = 0x80000009
	// MT 路由请求
	CMPP_MT_ROUTE       = 0x00000010
	// MT 路由请求应答
	CMPP_MT_ROUTE_RESP  = 0x80000010
	// MO 路由请求
	CMPP_MO_ROUTE       = 0x00000011
	// MO 路由请求应答
	CMPP_GET_ROUTE      = 0x00000012
	// 获取路由请求应答
	CMPP_GET_ROUTE_RESP = 0x80000012
	// MT 路由更新
	CMPP_MT_ROUTE_UPDATE = 0x00000013
	// MT 路由更新应答
	CMPP_MT_ROUTE_UPDATE_RESP = 0x80000013
	// MO 路由更新
	CMPP_MO_ROUTE_UPDATE = 0x00000014
	// MO 路由更新应答
	CMPP_MO_ROUTE_UPDATE_RESP = 0x80000014
	// MT 路由更新
	CMPP_PUSH_MT_ROUTE_UPDATE = 0x00000015
	// MT 路由更新应答
	CMPP_PUSH_MT_ROUTE_UPDATE_RESP = 0x80000015
	// MO 路由更新
	CMPP_PUSH_MO_ROUTE_UPDATE = 0x00000016
	// MO 路由更新应答
	CMPP_PUSH_MO_ROUTE_UPDATE_RESP = 0x80000016
)
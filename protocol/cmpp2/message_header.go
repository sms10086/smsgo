// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

// 7.3 消息头格式 （Message Header）
type MessageHeader struct {
	Total_Length uint32
	Command_Id uint32
	Sequence_Id uint32
}

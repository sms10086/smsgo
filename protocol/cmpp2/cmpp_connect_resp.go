// Copyright 2023 sms10086@hotmail.com. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

package cmpp2

type CmppConnectResp struct {
	Status uint8
	AuthenticatorISMG []byte
	Version uint8
}
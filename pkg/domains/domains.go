/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package domains

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// Matches returns one of the domains we're configured for
func Matches(s string) string { _ = "STUB: not implemented"; return "" }

// then we have a port and we just want to check the host

// IsUnderManagement check if an email is under vouch-managed domain
func IsUnderManagement(email string) bool { _ = "STUB: not implemented"; return false }

// ByLengthDesc sort from
// https://play.golang.org/p/N6GbEgBffd
type ByLengthDesc []string

func (s ByLengthDesc) Len() int { _ = "STUB: not implemented"; return 0 }

func (s ByLengthDesc) Swap(i, j int) { _ = "STUB: not implemented"; return }

// this differs by offing the longest first
func (s ByLengthDesc) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

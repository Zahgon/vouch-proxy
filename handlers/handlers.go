/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/vouch/vouch-proxy/pkg/structs"
)

// Provider each Provider must support GetuserInfo
type Provider interface {
	Configure()
	GetUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens, opts ...oauth2.AuthCodeOption) error
}

const (
	base64Bytes = 32
)

var (
	sessstore *sessions.CookieStore
	log       *zap.SugaredLogger
	fastlog   *zap.Logger
	provider  Provider
)

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// http://www.gorillatoolkit.org/pkg/sessions

// convert minutes to seconds

func getProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

// shouldn't ever reach this since cfg checks for a properly configure `oauth.provider`

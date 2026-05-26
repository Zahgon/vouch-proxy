/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package adfs

import (
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/vouch/vouch-proxy/pkg/structs"
)

// Provider provider specific functions
type Provider struct{}

type adfsTokenRes struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int64  `json:"expires_in"` // relative seconds from now
}

var log *zap.SugaredLogger

// Configure see main.go configure()
func (Provider) Configure() { _ = "STUB: not implemented"; return }

// GetUserInfo provider specific call to get userinfomation
// More info: https://docs.microsoft.com/en-us/windows-server/identity/ad-fs/overview/ad-fs-scenarios-for-developers#supported-scenarios
func (Provider) GetUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens, opts ...oauth2.AuthCodeOption) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// data contains an access token, refresh token, and id token
// Please note that in order for custom claims to work you MUST set allatclaims in ADFS to be passed
// https://oktotechnologies.ca/2018/08/26/adfs-openidconnect-configuration/

// If the email is blank, we will try to determine if the UPN is an email.

// Set the email from UPN if there is a valid email present.

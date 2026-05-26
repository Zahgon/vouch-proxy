/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package discord

import (
	"net/http"

	"golang.org/x/oauth2"

	"go.uber.org/zap"

	"github.com/vouch/vouch-proxy/pkg/structs"
)

// Provider provider specific functions
type Provider struct{}

var log *zap.SugaredLogger

// Configure see main.go configure()
func (Provider) Configure() { _ = "STUB: not implemented"; return }

// GetUserInfo provider specific call to get userinfomation
func (Provider) GetUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens, opts ...oauth2.AuthCodeOption) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// If the provider is configured to use IDs, the ID is copied to PreparedUsername.

// If the Discriminator is present that is appended to the Username in the format "Username#Discriminator"
// to match the old format of Discord usernames
// Previous format which is being phased out: https://support.discord.com/hc/en-us/articles/4407571667351-Law-Enforcement-Guidelines Subheading "How to find usernames and discriminators"
// Details about the new username requirements: https://support.discord.com/hc/en-us/articles/12620128861463

/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package jwtmanager

import (
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"github.com/vouch/vouch-proxy/pkg/structs"
)

const comma = ","

// VouchClaims jwt Claims specific to vouch
type VouchClaims struct {
	Username     string `json:"username"`
	CustomClaims map[string]interface{}
	PAccessToken string
	PIdToken     string
	jwt.RegisteredClaims
}

// RegisteredClaims jwt.RegisteredClaims implementation
var RegisteredClaims jwt.RegisteredClaims

var logger *zap.Logger
var log *zap.SugaredLogger
var aud []string

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// `aud` of the issued JWT https://tools.ietf.org/html/rfc7519#section-4.1.3
func audience() []string { _ = "STUB: not implemented"; return nil }

// TODO: the Sites that end up in the JWT come from here
// if we add fine grain ability (ACL?) to the equation
// then we're going to have to add something fancier here

// NewVPJWT issue a signed Vouch Proxy JWT for a user
func NewVPJWT(u structs.User, customClaims structs.CustomClaims, ptokens structs.PTokens) (string, error) {
	_ = "STUB: not implemented"
	// User`token`
	// u.PrepareUserData()
	return "", nil
}

// https://github.com/vouch/vouch-proxy/issues/287

// https://godoc.org/github.com/golang-jwt/jwt#NewWithClaims

// log.Debugf("token: %v", token)

// TODO: is this dead code?
// SiteInToken searches does the token contain the site?
func SiteInToken(site string, token *jwt.Token) bool { _ = "STUB: not implemented"; return false }

// ParseTokenString converts signed token to jwt struct
func ParseTokenString(tokenString string) (*jwt.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return jwt.ParseWithClaims(tokenString, &VouchClaims{}, func(token *jwt.Token) (interface{}, error) {

// SiteInAudience does the claim contain the value?
func (claims *VouchClaims) SiteInAudience(s string) bool { _ = "STUB: not implemented"; return false }

// jwt/v4 serialized the audience as a comma-separated string;
// jwt/v5 deserializes that into a single-element []string.

// PTokenClaims get all the claims
func PTokenClaims(ptoken *jwt.Token) (*VouchClaims, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeAndDecompressTokenString(encgzipss string) string {
	_ = "STUB: not implemented"

	// gzipss, err := url.QueryUnescape(encgzipss)
	return ""
}

func compressAndEncodeTokenString(ss string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ret := url.QueryEscape(buf.String())

// FindJWT look for JWT in Cookie, JWT Header, Authorization Header (OAuth2 Bearer Token)
// and Query String in that order
func FindJWT(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// ClaimsFromJWT parse the jwt and return the claims
func ClaimsFromJWT(jwt string) (*VouchClaims, error) { _ = "STUB: not implemented"; return nil, nil }

// claims = jwtmanager.PTokenClaims(jwtParsed)
// if claims == &jwtmanager.VouchClaims{} {

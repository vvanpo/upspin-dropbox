// Copyright 2017 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// The default Dropbox client configuration; this client uses the `Apps/upspin/`
// folder in the linked Dropbox account for storage.
var Config = &oauth2.Config{
	ClientID:     "wt1281n3q768jj3",
	ClientSecret: "blk944sx4oyf6aq",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	},
}

// Client returns an HTTP client that populates the Authorization header of all
// requests with an access token.
func Client(refreshToken string) *http.Client {
	token := oauth2.Token{RefreshToken: refreshToken}
	return Config.Client(context.Background(), &token)
}

// Exchange converts the authorization code into a refresh token.
func Exchange(authCode string) (string, error) {
	token, err := Config.Exchange(context.Background(), authCode)
	if err != nil {
		return "", err
	}

	return token.RefreshToken, nil
}

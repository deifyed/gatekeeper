package idtoken

import "github.com/coreos/go-oidc"

type Userinfoer interface {
	Extract(idToken string) (oidc.UserInfo, error)
}

type idTokenUserinfo struct {
	keySet   oidc.KeySet
	provider *oidc.Provider
}


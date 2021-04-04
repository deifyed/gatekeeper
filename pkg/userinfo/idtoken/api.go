package idtoken

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"time"
)

func New(issuer, jwksURL string) Userinfoer {
	ctx := context.Background()

	keySet := oidc.NewRemoteKeySet(ctx, jwksURL)
	
	provider, _ := oidc.NewProvider(ctx, issuer)
	
	return &idTokenUserinfo{
		keySet: keySet,
		provider: provider,
	}
}

type dummySource struct {
	token string
}

func (d dummySource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken:  "",
		TokenType:    "",
		RefreshToken: "",
		Expiry:       time.Time{},
	}, nil
}

func newDummySource(token string) oauth2.TokenSource {
	return &dummySource{token: token}
}

func (i idTokenUserinfo) Extract(idToken string) (oidc.UserInfo, error) {
	tokenSource := newDummySource(idToken)
	
	userinfo, err := i.provider.UserInfo(nil, tokenSource)
	if err != nil {
	    return oidc.UserInfo{}, fmt.Errorf("extracting userinfo: %w", err)
	}
	
	return *userinfo, nil
}

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
)

func main() {
	const (
		username   = ""
		password   = ""
		clientCode = ""
		partnerKey = ""
	)
	httpCli := common.GetDefaultHTTPClient()
	ctx := context.Background()
	session, err := auth.VerifyUserFull(ctx, username, password, clientCode, map[string]string{}, http.DefaultClient)
	if err != nil {
		panic(err)
	}
	sessInfo, err := auth.GetSessionKeyInfo(session.SessionKey, clientCode, httpCli)
	if err != nil {
		panic(err)
	}
	fmt.Println(sessInfo)

	info, err := auth.GetSessionKeyUser(session.SessionKey, clientCode, httpCli)
	cli, err := api.NewClient(session.SessionKey, clientCode, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	endpoints, err := cli.ServiceDiscoverer.GetServiceEndpoints(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(endpoints)

	partnerCli, err := api.NewPartnerClient(session.SessionKey, clientCode, partnerKey, nil)
	if err != nil {
		panic(err)
	}
	jwt, err := partnerCli.PartnerTokenProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}

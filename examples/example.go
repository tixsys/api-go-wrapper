package main

import (
	"context"
	"fmt"

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
	sessionKey, err := auth.VerifyUser(username, password, clientCode, httpCli)
	if err != nil {
		panic(err)
	}

	sessInfo, err := auth.GetSessionKeyInfo(sessionKey, clientCode, httpCli)
	if err != nil {
		panic(err)
	}
	fmt.Println(sessInfo)

	info, err := auth.GetSessionKeyUser(sessionKey, clientCode, httpCli)
	cli, err := api.NewClient(sessionKey, clientCode, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	endpoints, err := cli.ServiceDiscoverer.GetServiceEndpoints(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(endpoints)

	partnerCli, err := api.NewPartnerClient(sessionKey, clientCode, partnerKey, nil)
	if err != nil {
		panic(err)
	}
	jwt, err := partnerCli.PartnerTokenProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}

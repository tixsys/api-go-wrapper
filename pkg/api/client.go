package api

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

//IClient ...
type erplyClient struct {
	sessionKey string
	clientCode string
	partnerKey string
	secret     string
	url        string
	httpClient *http.Client
}

//VerifyUser will give you session key
func VerifyUser(username string, password string, clientCode string, client *http.Client) (string, error) {
	requestUrl := fmt.Sprintf(baseURL, clientCode)
	params := url.Values{}
	params.Add("username", username)
	params.Add("clientCode", clientCode)
	params.Add("password", password)
	params.Add("request", "verifyUser")

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return "", erplyerr("failed to build HTTP request", err)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", erplyerr("failed to build VerifyUser request", err)
	}

	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", erplyerr("failed to decode VerifyUserResponse", err)
	}
	if len(res.Records) < 1 {
		return "", erplyerr("VerifyUser: no records in response", nil)
	}
	return res.Records[0].SessionKey, nil
}

// NewClient Takes three params:
// sessionKey string obtained from credentials or jwt
// clientCode erply customer identification number
// and a custom http Client if needs to be overwritten. if nil will use default http client provided by the SDK
func NewClient(sessionKey string, clientCode string, customCli *http.Client) IClient {

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 4 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,

		MaxIdleConns:    MaxIdleConns,
		MaxConnsPerHost: MaxConnsPerHost,
	}

	cli := erplyClient{
		sessionKey: sessionKey,
		clientCode: clientCode,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: &http.Client{
			Transport: tr,
			Timeout:   5 * time.Second,
		},
	}
	if customCli != nil {
		cli.httpClient = customCli
	}
	return &cli
}

func NewClientV2(partnerKey string, secret string, clientCode string) IClient {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 4 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,

		MaxIdleConns:    MaxIdleConns,
		MaxConnsPerHost: MaxConnsPerHost,
	}

	cli := erplyClient{
		partnerKey: partnerKey,
		secret:     secret,
		clientCode: clientCode,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: &http.Client{
			Transport: tr,
			Timeout:   5 * time.Second,
		},
	}
	return &cli
}
func (cli *erplyClient) Close() {
	cli.httpClient.CloseIdleConnections()
}

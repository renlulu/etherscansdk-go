package internal

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type Contract struct {
	Base
}

func NewContract(api string, apiKey string) *Contract {
	client := resty.New()
	return &Contract{
		Base{
			API:    api,
			APIKey: apiKey,
			Client: client,
		},
	}
}

func (c *Contract) CheckProxyVerification(ctx context.Context, guid string) (string, error) {
	resp, err := c.Client.R().
		SetQueryParam("module", "contract").
		SetQueryParam("action", "checkproxyverification").
		SetQueryParam("apikey", c.APIKey).
		SetQueryParam("guid", guid).
		Get(c.API)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (c *Contract) VerifyProxy(
	ctx context.Context,
	address string,
) (string, error) {
	u := fmt.Sprintf(
		"%s?module=contract&action=verifyproxycontract&apikey=%s&address=%s",
		c.API, c.APIKey, address,
	)
	requestData := fmt.Sprintf("address=%s", address)
	fmt.Println(requestData)
	resp, err := c.Client.
		R().SetBody(requestData).
		//SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(u)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (c *Contract) VerifySourceCode(
	ctx context.Context,
	chainId string,
	address string,
	sourceCode string,
	arguments string,
	contractName string,
	compilerVersion string,
) (string, error) {
	requestData := fmt.Sprintf(
		"module=contract&action=verifysourcecode&apikey=%s&chainId=%s&codeformat=solidity-standard-json-input&sourceCode=%s&constructorArguements=%s&contractaddress=%s&contractname=%s&compilerversion=%s",
		c.APIKey, chainId, url.QueryEscape(sourceCode), url.QueryEscape(arguments), address, contractName, url.QueryEscape(compilerVersion),
	)
	resp, err := c.Client.
		R().SetBody(requestData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(c.API)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (c *Contract) GetABI(address string) (string, error) {
	resp, err := c.Client.
		R().
		SetQueryParam("module", "contract").
		SetQueryParam("action", "getabi").
		SetQueryParam("apikey", c.APIKey).
		SetQueryParam("address", address).
		Post(c.API)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

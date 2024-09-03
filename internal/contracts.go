package internal

import "github.com/go-resty/resty/v2"

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

func (c *Contract) GetContractABI(address string) (string, error) {
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

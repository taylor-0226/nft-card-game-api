package venly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	VENLY_API_URL    = "https://api-wallet-sandbox.venly.io"
	VENLY_AUTH_URL   = "https://login-sandbox.venly.io"
	VENLY_CLIENT_ID  = "3e572e5a-8c4d-466f-a674-02a030f9aa57"
	VENLY_APP_SECRET = "2971ea69-564e-4c71-833d-f29d9217495b"
)

var Global *VenlyClient

type VenlyRequestCreateWallet struct {
	Description string `json:"description,omitempty"`
	PinCode     string `json:"pincode,omitempty"`
	SecretType  string `json:"secretType,omitempty"`
	WalletType  string `json:"walletType,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}

type VenlyWallet struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	WalletType   string `json:"walletType"`
	SecretType   string `json:"secretType"`
	CreatedAt    string `json:"createdAt"`
	Archived     bool   `json:"archived"`
	Description  string `json:"description"`
	Primary      bool   `json:"primary"`
	HasCustomPin bool   `json:"hasCustomPin"`
	Identifier   string `json:"identifier"`
	Balance      struct {
		Available     bool   `json:"available"`
		SecretType    string `json:"secretType"`
		Balance       int    `json:"balance"`
		GasBalance    int    `json:"gasBalance"`
		Symbol        string `json:"symbol"`
		GasSymbol     string `json:"gasSymbol"`
		RawBalance    string `json:"rawBalance"`
		RawGasBalance string `json:"rawGasBalance"`
		Decimals      int    `json:"decimals"`
	} `json:"balance"`
}

type VenlyResponseCreateWallet struct {
	Success bool        `json:"success"`
	Result  VenlyWallet `json:"result"`
}

type VenlyRequestGetAccessToken struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type VenlyResponseGetAccessToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type VenlyClient struct {
	Http        http.Client
	Config      VenlyClientConfig
	AccessToken string
	ExpiresAt   time.Time
}

type VenlyClientConfig struct {
	ClientId     string
	ClientSecret string
}

func NewClient(config VenlyClientConfig) (*VenlyClient, error) {
	client := VenlyClient{
		Config: config,
		Http:   *http.DefaultClient,
	}

	return &client, nil
}

func (c *VenlyClient) SendRequest(t string, url string, data interface{}) ([]byte, error) {
	if time.Now().After(c.ExpiresAt) {
		err := c.GetAccessToken()
		if err != nil {
			return nil, err
		}
	}

	if t == "POST" {
		if data != nil {
			js, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}

			r, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
			if err != nil {
				return nil, err
			}

			r.Header.Add("Content-Type", "application/json")
			r.Header.Add("Authorization", "Bearer "+c.AccessToken)

			res, err := c.Http.Do(r)
			if err != nil {
				return nil, err
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}

			fmt.Println("Venly send request response: ", string(b))

			if res.StatusCode != 200 {
				return nil, errors.New(fmt.Sprintf("Failed POST request to Venly %s: %s", url, res.Status))
			}

			return b, nil
		}

		r, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return nil, err
		}

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", "Bearer "+c.AccessToken)

		res, err := c.Http.Do(r)
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 200 {
			return nil, errors.New(fmt.Sprintf("Failed POST request to Venly %s: %s", url, res.Status))
		}

		return b, nil
	}

	if t == "GET" {
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", "Bearer "+c.AccessToken)

		res, err := c.Http.Do(r)
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 200 {
			return nil, errors.New(fmt.Sprintf("Failed GET request to Venly %s: %s", url, res.Status))
		}

		return b, nil

	}

	return nil, errors.New("Invalid request type")
}

func (c *VenlyClient) GetAccessToken() error {
	uri := VENLY_AUTH_URL + "/auth/realms/Arkane/protocol/openid-connect/token"

	data := url.Values{
		"client_id":     {c.Config.ClientId},
		"client_secret": {c.Config.ClientSecret},
		"grant_type":    {"client_credentials"},
	}

	r, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Http.Do(r)
	if err != nil {
		return err
	}

	var resp VenlyResponseGetAccessToken
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		return err
	}

	fmt.Println("Venly get access token response: ", string(b))

	if res.StatusCode != 200 {
		return errors.New(res.Status + ": failed to get access token")
	}

	c.AccessToken = resp.AccessToken
	c.ExpiresAt = time.Now().Add(time.Duration(resp.ExpiresIn) * time.Second)

	return nil
}

func (c *VenlyClient) CreateWallet(data VenlyRequestCreateWallet) (VenlyWallet, error) {
	var wallet VenlyWallet
	url := VENLY_API_URL + "/api/wallets"

	// data := VenlyRequestCreateWallet{
	// 	Description: "test",
	// 	PinCode:     "1234",
	// 	SecretType:  "MATIC",
	// 	WalletType:  "WHITE_LABEL",
	// 	Identifier:  "type=unrecoverable",
	// }

	var resp VenlyResponseCreateWallet

	b, err := c.SendRequest("POST", url, data)
	if err != nil {
		return wallet, err
	}

	fmt.Println("Venly create wallet response: ", string(b))

	err = json.Unmarshal(b, &resp)
	if err != nil {
		return wallet, err
	}

	fmt.Println("Venly create wallet: ", resp)

	if !resp.Success {
		return wallet, errors.New("failed to create wallet")
	}

	wallet = resp.Result

	return wallet, nil
}

type VenlyRequestTransferNft struct {
	Pincode            int `json:"pincode"`
	TransactionRequest struct {
		Type         string `json:"type"`
		SecretType   string `json:"secretType"`
		WalletID     string `json:"walletId"`
		To           string `json:"to"`
		TokenAddress string `json:"tokenAddress"`
		TokenID      int    `json:"tokenId"`
	} `json:"transactionRequest"`
}

type VenlyResponseTransferNft struct {
}

// func (c *VenlyClient) TransferNft(data VenlyRequestTransferNft) (VenlyResponseTransferNft, error) {
// 	var res VenlyResponseTransferNft
// 	url := VENLY_API_URL + "/api/transactions/execute"

// 	// {
// 	// 	"pincode": 1234,
// 	// 	"transactionRequest": {
// 	// 		"type": "NFT_TRANSFER",
// 	// 		"secretType": "MATIC",
// 	// 		"walletId": "89733529-ddc1-44aa-b2af-17601e2e26c0",
// 	// 		"to": "0x9d9376EEbFE3443544d3654f7aD272f0A31D8152",
// 	// 		"tokenAddress": "0xd05a795d339886bb8dd46cfe2ac009d7f1e48a64",
// 	// 		"tokenId": 91
// 	// 	}
// 	// }
// 	return res, nil
// }

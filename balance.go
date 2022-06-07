package digiflazz

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Balance struct {
	API *API
}

type BalanceBadRequest struct {
	Data struct {
		ResponseCode string `json:"rc"`
		Depsoit      uint   `json:"deposit"`
		Mesaage      string `json:"message"`
	} `json:"data"`
}

type BalanceResponse struct {
	Data struct {
		Saldo float64 `json:"saldo"`
	} `json:"data"`
}

type BalanceRequest struct {
	Cmd      string    `json:"cmd"`
	Username string    `json:"username"`
	Sign     SignedKey `json:"sign"`
}

func (c *Balance) CekSaldo() (float64, error) {
	var balanceResponse BalanceResponse

	endpoint := fmt.Sprintf("%s/%s/cek-saldo", c.API.BaseURL, c.API.APIVersion)
	reqData := BalanceRequest{
		Cmd:      "deposit",
		Username: c.API.opt.Username,
		Sign:     SignMD5(c.API.opt.Username, c.API.opt.Key, "depo"),
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	res, err := CallAPIRequest("POST", endpoint, header, reqData, balanceResponse)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&balanceResponse); err != nil {
		return 0, err
	}

	return balanceResponse.Data.Saldo, nil
}

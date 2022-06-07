package digiflazz

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Deposit struct {
	API *API
}

type Bank string

const (
	BankMandiri Bank = "MANDIRI"
	BankBCA     Bank = "BCA"
	BankBRI     Bank = "BRI"
)

type DepositRequest struct {
	Username  string    `json:"username"`
	Amount    uint      `json:"amount"`
	Bank      Bank      `json:"bank"`
	OwnerName string    `json:"owner_name"`
	Sign      SignedKey `json:"sign"`
}

type DepositResponseData struct {
	ResponseCode string `json:"rc"`
	Amount       uint   `json:"amount"`
	Notes        string `json:"notes"`
}

type DepositResponse struct {
	Data DepositResponseData `json:"data"`
}

func (c *Deposit) BuatTiket(amount uint, bank Bank, ownerName string) (DepositResponseData, error) {
	if amount < 200000 {
		return DepositResponseData{}, errors.New("Amount minimal 200000")
	}

	var depositResponse DepositResponse

	endpoint := fmt.Sprintf("%s/%s/deposit", c.API.BaseURL, c.API.APIVersion)
	reqData := DepositRequest{
		Username:  c.API.opt.Username,
		Amount:    amount,
		Bank:      bank,
		OwnerName: ownerName,
		Sign:      SignMD5(c.API.opt.Username, c.API.opt.Key, "deposit"),
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	res, err := CallAPIRequest("POST", endpoint, header, reqData, depositResponse)
	if err != nil {
		return DepositResponseData{}, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&depositResponse); err != nil {
		return DepositResponseData{}, err
	}

	return depositResponse.Data, nil
}

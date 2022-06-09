package digiflazz

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Transaction struct {
	API *API
}

type TopupRequest struct {
	Username       string    `json:"username"`
	BuyerSkuCode   string    `json:"buyer_sku_code"`
	CustomerNumber string    `json:"customer_no"`
	RefrenceId     string    `json:"ref_id"`
	Sign           SignedKey `json:"sign"`
	Testing        bool      `json:"testing"`
	Message        string    `json:"message"`
}

type TopupResponse struct {
	RefrenceId     string  `json:"ref_id"`
	CustomerNumber string  `json:"customer_no"`
	BuyerSkuCode   string  `json:"buyer_sku_code"`
	Message        string  `json:"message"`
	Status         string  `json:"status"`
	ResponseCode   string  `json:"response_code"`
	SerialNumber   string  `json:"sn,omitempty"`
	BuyerLastSaldo float64 `json:"buyer_last_saldo,omitempty"`
	Price          uint    `json:"price"`
	Telegram       string  `json:"telegram,omitempty"`
	Whatsapp       string  `json:"whatsapp,omitempty"`
}

type ResponseTopup struct {
	Data TopupResponse `json:"data"`
}

func (c *Transaction) Topup(buyerSkuCode string, customerNumber string, refrenceId string, testing bool, message string) (TopupResponse, error) {
	var responseTopup ResponseTopup

	endpoint := fmt.Sprintf("%s/%s/transaction", c.API.BaseURL, c.API.APIVersion)
	reqData := TopupRequest{
		Username:       c.API.opt.Username,
		BuyerSkuCode:   buyerSkuCode,
		CustomerNumber: customerNumber,
		RefrenceId:     refrenceId,
		Sign:           SignMD5(c.API.opt.Username, c.API.opt.Key, refrenceId),
		Testing:        testing,
		Message:        message,
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	res, err := CallAPIRequest("POST", endpoint, header, reqData, responseTopup)
	if err != nil {
		return TopupResponse{}, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&responseTopup); err != nil {
		return TopupResponse{}, err
	}

	return responseTopup.Data, nil
}

// Refer to: https://developer.digiflazz.com/api/buyer/cek-status/#prepaid
func (c *Transaction) CekTopup(refrenceId string) (TopupResponse, error) {
	var responseTopup ResponseTopup

	endpoint := fmt.Sprintf("%s/%s/transaction", c.API.BaseURL, c.API.APIVersion)
	reqData := TopupRequest{
		Username:   c.API.opt.Username,
		RefrenceId: refrenceId,
		Sign:       SignMD5(c.API.opt.Username, c.API.opt.Key, refrenceId),
		Testing:    false,
		Message:    "",
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	res, err := CallAPIRequest("POST", endpoint, header, reqData, responseTopup)
	if err != nil {
		return TopupResponse{}, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&responseTopup); err != nil {
		return TopupResponse{}, err
	}

	return responseTopup.Data, nil
}

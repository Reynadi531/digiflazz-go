package digiflazz

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pricelist struct {
	API *API
}

type Product struct {
	ProductName         string `json:"product_name"`
	Category            string `json:"category"`
	Brand               string `json:"brand"`
	Type                string `json:"type"`
	SellerName          string `json:"seller_name"`
	Price               uint   `json:"price"`
	BuyerSkuCode        string `json:"buyer_sku_code"`
	BuyerProductStatus  bool   `json:"buyer_product_status"`
	SellerProductStatus bool   `json:"seller_product_status"`
	UnlimitedStock      bool   `json:"unlimited_stock"`
	Stock               uint   `json:"stock"`
	Multi               bool   `json:"multi"`
	StartCutOff         string `json:"start_cut_off"`
	EndCutOff           string `json:"end_cut_off"`
	Description         string `json:"desc"`
}

type PricelistRequest struct {
	Cmd      string    `json:"cmd"`
	Username string    `json:"username"`
	Code     string    `json:"code"`
	Sign     SignedKey `json:"sign"`
}

type pricelistResponse struct {
	Data []Product `json:"data"`
}

func (c *Pricelist) CekHarga(code ...string) ([]Product, error) {
	var pricelistResponse pricelistResponse
	var reqData PricelistRequest

	endpoint := fmt.Sprintf("%s/%s/price-list", c.API.BaseURL, c.API.APIVersion)
	reqData = PricelistRequest{
		Cmd:      "price-list",
		Username: c.API.opt.Username,
		Code:     "",
		Sign:     SignMD5(c.API.opt.Username, c.API.opt.Key, "pricelist"),
	}

	if len(code) > 0 {
		reqData.Code = code[0]
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	res, err := CallAPIRequest("POST", endpoint, header, reqData, pricelistResponse)
	if err != nil {
		return pricelistResponse.Data, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&pricelistResponse); err != nil {
		return pricelistResponse.Data, err
	}

	return pricelistResponse.Data, nil
}

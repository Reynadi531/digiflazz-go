package digiflazz

import "net/http"

func init() {

}

func New(username string, key string) *API {
	var api API

	api.opt.Username = username
	api.opt.Key = key
	api.APIVersion = ApiVersionV1
	api.HttpClient = &http.Client{}
	api.BaseURL = "https://api.digiflazz.com"

	api.Balance.API = &api
	api.Pricelist.API = &api
	api.Deposit.API = &api
	api.Transaction.API = &api

	return &api
}

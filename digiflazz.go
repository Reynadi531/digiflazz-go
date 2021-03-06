package digiflazz

import "net/http"

var api API

func init() {
	api.Balance.API = &api
	api.Pricelist.API = &api
	api.Deposit.API = &api
	api.Transaction.API = &api
}

func New(username string, key string) *API {
	api.opt.Username = username
	api.opt.Key = key
	api.APIVersion = ApiVersionV1
	api.HttpClient = &http.Client{}
	api.BaseURL = "https://api.digiflazz.com"

	return &api
}

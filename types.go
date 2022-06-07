package digiflazz

import "net/http"

type SignedKey string
type APIVersion string

const (
	ApiVersionV1 APIVersion = "v1"
)

type Option struct {
	Username string
	Key      string
}

type API struct {
	BaseURL     string
	opt         Option
	Balance     Balance
	Pricelist   Pricelist
	Deposit     Deposit
	Transaction Transaction
	APIVersion  APIVersion
	SignedKey   SignedKey
	HttpClient  *http.Client
}

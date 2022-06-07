package digiflazz

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func CallAPIRequest(method string, url string, header http.Header, body interface{}, result interface{}) (*http.Response, error) {
	reqBody := []byte("")
	var err error
	var req *http.Request
	var client http.Client

	reqBody, err = json.Marshal(body)
	if err != nil {
		return nil, errors.New("Error marshal request data :" + err.Error())
	}

	req, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if header != nil {
		req.Header = header
	}
	req.Header.Set("User-Agent", "digiflazz-go")

	if err != nil {
		return nil, errors.New("Failed creating request :" + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed HTTP request : " + err.Error())
	}

	return res, err
}

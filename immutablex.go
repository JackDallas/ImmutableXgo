package immutablex

import (
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL = "https://api.x.immutable.com/v1/"
)

type Client struct {
}

func (c *Client) doRequest(method, url string, body interface{}) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

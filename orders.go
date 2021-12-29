package immutablex

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

func (c *Client) OrdersGET(props *OrdersGETProps) (OrdersGETResponse, error) {

	url := BASE_URL + "orders"

	if props != nil {
		v, err := query.Values(props)
		if err != nil {
			return OrdersGETResponse{}, err
		}
		url = url + "?" + v.Encode()
	}

	b, err := c.doRequest("GET", url, nil)
	if err != nil {
		return OrdersGETResponse{}, err
	}
	var resp OrdersGETResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return OrdersGETResponse{}, err
	}
	return resp, nil
}

type OrdersGETResponse struct {
	Cursor    string `json:"cursor"`
	Remaining int    `json:"remaining"`
	Result    []struct {
		AmountSold string `json:"amount_sold"`
		Buy        struct {
			Data struct {
				Decimals   int    `json:"decimals"`
				ID         string `json:"id"`
				Properties struct {
					Collection struct {
						IconURL string `json:"icon_url"`
						Name    string `json:"name"`
					} `json:"collection"`
					ImageURL string `json:"image_url"`
					Name     string `json:"name"`
				} `json:"properties"`
				Quantity     string `json:"quantity"`
				TokenAddress string `json:"token_address"`
				TokenID      string `json:"token_id"`
			} `json:"data"`
			Type string `json:"type"`
		} `json:"buy"`
		ExpirationTimestamp string `json:"expiration_timestamp"`
		Fees                []struct {
			Address string `json:"address"`
			Amount  string `json:"amount"`
			Token   struct {
				Data struct {
					ContractAddress string `json:"contract_address"`
					Decimals        int    `json:"decimals"`
				} `json:"data"`
				Type string `json:"type"`
			} `json:"token"`
			Type string `json:"type"`
		} `json:"fees"`
		OrderID int `json:"order_id"`
		Sell    struct {
			Data struct {
				Decimals   int    `json:"decimals"`
				ID         string `json:"id"`
				Properties struct {
					Collection struct {
						IconURL string `json:"icon_url"`
						Name    string `json:"name"`
					} `json:"collection"`
					ImageURL string `json:"image_url"`
					Name     string `json:"name"`
				} `json:"properties"`
				Quantity     string `json:"quantity"`
				TokenAddress string `json:"token_address"`
				TokenID      string `json:"token_id"`
			} `json:"data"`
			Type string `json:"type"`
		} `json:"sell"`
		Status           string `json:"status"`
		Timestamp        string `json:"timestamp"`
		UpdatedTimestamp string `json:"updated_timestamp"`
		User             string `json:"user"`
	} `json:"result"`
}
type OrdersGETProps struct {
	//Page size of the result
	PageSize int `url:"page_size,omitempty"`
	//Cursor
	Cursor string `url:"cursor,omitempty"`
	//Property to sort by
	OrderBy string `url:"order_by,omitempty"`
	//Direction to sort (asc/desc)
	Direction string `url:"direction,omitempty"`
	//Include network fees
	IncludeFees bool `url:"include_fees,omitempty"`
	//Ethereum address of the user who submitted this order
	User string `url:"user,omitempty"`
	//Status of this order
	Status string `url:"status,omitempty"`
	//Minimum created at timestamp for this order
	MinTimestamp string `url:"min_timestamp,omitempty"`
	//Maximum created at timestamp for this order
	MaxTimestamp string `url:"max_timestamp,omitempty"`
	//Minimum updated at timestamp for this order
	UpdatedMinTimestamp string `url:"updated_min_timestamp,omitempty"`
	//Maximum updated at timestamp for this order
	UpdatedMaxTimestamp string `url:"updated_max_timestamp,omitempty"`
	//Token type of the asset this order buys
	BuyTokenType string `url:"buy_token_type,omitempty"`
	//ERC721 Token ID of the asset this order buys
	BuyTokenID string `url:"buy_token_id,omitempty"`
	//Internal IMX ID of the asset this order buys
	BuyAssetID string `url:"buy_asset_id,omitempty"`
	//Comma separated string of token addresses of the asset this order buys
	BuyTokenAddress string `url:"buy_token_address,omitempty"`
	//Token name of the asset this order buys
	BuyTokenName string `url:"buy_token_name,omitempty"`
	//Min quantity for the asset this order buys
	BuyMinQuantity string `url:"buy_min_quantity,omitempty"`
	//Max quantity for the asset this order buys
	BuyMaxQuantity string `url:"buy_max_quantity,omitempty"`
	//JSON-encoded metadata filters for the asset this order buys
	BuyMetadata string `url:"buy_metadata,omitempty"`
	//Token type of the asset this order sells
	SellTokenType string `url:"sell_token_type,omitempty"`
	//ERC721 Token ID of the asset this order sells
	SellTokenID string `url:"sell_token_id,omitempty"`
	//Internal IMX ID of the asset this order sells
	SellAssetID string `url:"sell_asset_id,omitempty"`
	//Comma separated string of token addresses of the asset this order sells
	SellTokenAddress string `url:"sell_token_address,omitempty"`
	//Token name of the asset this order sells
	SellTokenName string `url:"sell_token_name,omitempty"`
	//Min quantity for the asset this order sells
	SellMinQuantity string `url:"sell_min_quantity,omitempty"`
	//Max quantity for the asset this order sells
	SellMaxQuantity string `url:"sell_max_quantity,omitempty"`
	//JSON-encoded metadata filters for the asset this order sells
	SellMetadata string `url:"sell_metadata,omitempty"`
}

package models

import "encoding/json"

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data json.RawMessage `json:"data"`
}

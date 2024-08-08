package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data json.RawMessage `json:"data"` //represt the respons data from the API
}

type CreateAlertPolicyResponse struct {
	AlertsPolicyCreate struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"alertsPolicyCreate"`
}

func CreateAlertPolicy(name string) (string, error) {
	query := fmt.Sprintf(`
		mutation {
			alertsPolicyCreate(
				accountId: %s,
				policy: {
					name: "%s",
					incidentPreference: PER_POLICY
				}
			) {
				id
				name
			}
		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), name)

	response, err := makeGraphQLRequest(query)
	if err != nil {
		return "", err
	}

	var result CreateAlertPolicyResponse
	err = json.Unmarshal(response.Data, &result)
	if err != nil {
		return "", err
	}

	return result.AlertsPolicyCreate.ID, nil
}

func UpdateAlertPolicy(id, name string) error {
	query := fmt.Sprintf(`
		mutation {
			alertsPolicyUpdate(
				accountId: %s,
				id: "%s",
				policy: {
					name: "%s",
					incidentPreference: PER_POLICY
				}
			) {
				id
				name
			}
		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), id, name)

	_, err := makeGraphQLRequest(query)
	return err
}

func DeleteAlertPolicy(id string) error {
	query := fmt.Sprintf(`
		mutation {
			alertsPolicyDelete(
				accountId: %s,
				id: "%s"
			) {
				id
			}
		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), id)

	_, err := makeGraphQLRequest(query)
	return err
}

func FetchAlertPolicy(id string) ([]byte, error) {
	query := fmt.Sprintf(`
		{
			actor {
				account(id: %s) {
					alerts {
						policy(id: "%s") {
							id
							name
							incidentPreference
						}
					}
				}
			}
		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), id)

	response, err := makeGraphQLRequest(query)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func makeGraphQLRequest(query string) (GraphQLResponse, error) {
	url := os.Getenv("NEW_RELIC_API_URL")
	apiKey := os.Getenv("NEW_RELIC_API_KEY")

	requestBody, err := json.Marshal(GraphQLRequest{Query: query})
	if err != nil {
		return GraphQLResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return GraphQLResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GraphQLResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GraphQLResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GraphQLResponse{}, fmt.Errorf("API error: %s", string(body))
	}

	var response GraphQLResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

// package services

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// type GraphQLRequest struct {
// 	Query string `json:"query"`
// }

// type GraphQLResponse struct {
// 	Data json.RawMessage `json:"data"`
// }

// func CreateAlertPolicy(name string) (string, error) {
// 	query := fmt.Sprintf(`
// 		mutation {
// 			alertsPolicyCreate(
// 				accountId: %s,
// 				policy: {
// 					name: "%s",
// 					incidentPreference: PER_POLICY
// 				}
// 			) {
// 				id
// 				name
// 			}
// 		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), name)

// 	response, err := makeGraphQLRequest(query)
// 	if err != nil {
// 		return "", err
// 	}

// 	var result map[string]map[string]map[string]string
// 	err = json.Unmarshal(response.Data, &result)
// 	if err != nil {
// 		return "", err
// 	}

// 	return result["alertsPolicyCreate"]["id"], nil
// }

// func UpdateAlertPolicy(id, name string) error {
// 	query := fmt.Sprintf(`
// 		mutation {
// 			alertsPolicyUpdate(
// 				accountId: %s,
// 				id: "%s",
// 				policy: {
// 					name: "%s",
// 					incidentPreference: PER_POLICY
// 				}
// 			) {
// 				id
// 				name
// 			}
// 		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), id, name)

// 	_, err := makeGraphQLRequest(query)
// 	return err
// }

// func DeleteAlertPolicy(id string) error {
// 	query := fmt.Sprintf(`
// 		mutation {
// 			alertsPolicyDelete(
// 				accountId: %s,
// 				id: "%s"
// 			) {
// 				id
// 			}
// 		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), id)

// 	_, err := makeGraphQLRequest(query)
// 	return err
// }

// func FetchAlertPolicy(id string) ([]byte, error) {
// 	query := fmt.Sprintf(`
// 		{
// 			actor {
// 				account(id: %s) {
// 					alerts {
// 						policy(id: "%s") {
// 							id
// 							name
// 							incidentPreference
// 						}
// 					}
// 				}
// 			}
// 		}`, os.Getenv("NEW_RELIC_ACCOUNT_ID"), id)

// 	response, err := makeGraphQLRequest(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response.Data, nil
// }

// func makeGraphQLRequest(query string) (GraphQLResponse, error) {
// 	url := os.Getenv("NEW_RELIC_API_URL")
// 	apiKey := os.Getenv("NEW_RELIC_API_KEY")

// 	requestBody, err := json.Marshal(GraphQLRequest{Query: query})
// 	if err != nil {
// 		return GraphQLResponse{}, err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		return GraphQLResponse{}, err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Api-Key", apiKey)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return GraphQLResponse{}, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return GraphQLResponse{}, err
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return GraphQLResponse{}, fmt.Errorf("API error: %s", string(body))
// 	}

// 	var response GraphQLResponse
// 	err = json.Unmarshal(body, &response)
// 	return response, err
// }

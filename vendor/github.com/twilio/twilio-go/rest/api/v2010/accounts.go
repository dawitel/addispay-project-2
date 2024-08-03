/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateAccount'
type CreateAccountParams struct {
	// A human readable description of the account to create, defaults to `SubAccount Created at {YYYY-MM-DD HH:MM meridian}`
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateAccountParams) SetFriendlyName(FriendlyName string) *CreateAccountParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new Twilio Subaccount from the account making the request
func (c *ApiService) CreateAccount(params *CreateAccountParams) (*ApiV2010Account, error) {
	path := "/2010-04-01/Accounts.json"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Account{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch the account specified by the provided Account Sid
func (c *ApiService) FetchAccount(Sid string) (*ApiV2010Account, error) {
	path := "/2010-04-01/Accounts/{Sid}.json"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Account{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAccount'
type ListAccountParams struct {
	// Only return the Account resources with friendly names that exactly match this name.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Only return Account resources with the given status. Can be `closed`, `suspended` or `active`.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAccountParams) SetFriendlyName(FriendlyName string) *ListAccountParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListAccountParams) SetStatus(Status string) *ListAccountParams {
	params.Status = &Status
	return params
}
func (params *ListAccountParams) SetPageSize(PageSize int) *ListAccountParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAccountParams) SetLimit(Limit int) *ListAccountParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Account records from the API. Request is executed immediately.
func (c *ApiService) PageAccount(params *ListAccountParams, pageToken, pageNumber string) (*ListAccountResponse, error) {
	path := "/2010-04-01/Accounts.json"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAccountResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Account records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAccount(params *ListAccountParams) ([]ApiV2010Account, error) {
	response, errors := c.StreamAccount(params)

	records := make([]ApiV2010Account, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Account records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAccount(params *ListAccountParams) (chan ApiV2010Account, chan error) {
	if params == nil {
		params = &ListAccountParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010Account, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAccount(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAccount(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAccount(response *ListAccountResponse, params *ListAccountParams, recordChannel chan ApiV2010Account, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Accounts
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAccountResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAccountResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAccountResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAccountResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateAccount'
type UpdateAccountParams struct {
	// Update the human-readable description of this Account
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateAccountParams) SetFriendlyName(FriendlyName string) *UpdateAccountParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateAccountParams) SetStatus(Status string) *UpdateAccountParams {
	params.Status = &Status
	return params
}

// Modify the properties of a given Account
func (c *ApiService) UpdateAccount(Sid string, params *UpdateAccountParams) (*ApiV2010Account, error) {
	path := "/2010-04-01/Accounts/{Sid}.json"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Account{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
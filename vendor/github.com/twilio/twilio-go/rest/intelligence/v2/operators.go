/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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

// Fetch a specific Operator. Works for both Custom and Pre-built Operators.
func (c *ApiService) FetchOperator(Sid string) (*IntelligenceV2Operator, error) {
	path := "/v2/Operators/{Sid}"
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

	ps := &IntelligenceV2Operator{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListOperator'
type ListOperatorParams struct {
	// Returns Operators with the provided availability type. Possible values: internal, beta, public, retired.
	Availability *string `json:"Availability,omitempty"`
	// Returns Operators that support the provided language code.
	LanguageCode *string `json:"LanguageCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListOperatorParams) SetAvailability(Availability string) *ListOperatorParams {
	params.Availability = &Availability
	return params
}
func (params *ListOperatorParams) SetLanguageCode(LanguageCode string) *ListOperatorParams {
	params.LanguageCode = &LanguageCode
	return params
}
func (params *ListOperatorParams) SetPageSize(PageSize int) *ListOperatorParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListOperatorParams) SetLimit(Limit int) *ListOperatorParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Operator records from the API. Request is executed immediately.
func (c *ApiService) PageOperator(params *ListOperatorParams, pageToken, pageNumber string) (*ListOperatorResponse, error) {
	path := "/v2/Operators"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Availability != nil {
		data.Set("Availability", *params.Availability)
	}
	if params != nil && params.LanguageCode != nil {
		data.Set("LanguageCode", *params.LanguageCode)
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

	ps := &ListOperatorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Operator records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListOperator(params *ListOperatorParams) ([]IntelligenceV2Operator, error) {
	response, errors := c.StreamOperator(params)

	records := make([]IntelligenceV2Operator, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Operator records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamOperator(params *ListOperatorParams) (chan IntelligenceV2Operator, chan error) {
	if params == nil {
		params = &ListOperatorParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IntelligenceV2Operator, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageOperator(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamOperator(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamOperator(response *ListOperatorResponse, params *ListOperatorParams, recordChannel chan IntelligenceV2Operator, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Operators
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListOperatorResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListOperatorResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListOperatorResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListOperatorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

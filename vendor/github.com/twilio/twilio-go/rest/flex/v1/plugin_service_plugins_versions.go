/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// Optional parameters for the method 'CreatePluginVersion'
type CreatePluginVersionParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
	// The Flex Plugin Version's version.
	Version *string `json:"Version,omitempty"`
	// The URL of the Flex Plugin Version bundle
	PluginUrl *string `json:"PluginUrl,omitempty"`
	// The changelog of the Flex Plugin Version.
	Changelog *string `json:"Changelog,omitempty"`
	// Whether this Flex Plugin Version requires authorization.
	Private *bool `json:"Private,omitempty"`
	// The version of Flex Plugins CLI used to create this plugin
	CliVersion *string `json:"CliVersion,omitempty"`
	// The validation status of the plugin, indicating whether it has been validated
	ValidateStatus *string `json:"ValidateStatus,omitempty"`
}

func (params *CreatePluginVersionParams) SetFlexMetadata(FlexMetadata string) *CreatePluginVersionParams {
	params.FlexMetadata = &FlexMetadata
	return params
}
func (params *CreatePluginVersionParams) SetVersion(Version string) *CreatePluginVersionParams {
	params.Version = &Version
	return params
}
func (params *CreatePluginVersionParams) SetPluginUrl(PluginUrl string) *CreatePluginVersionParams {
	params.PluginUrl = &PluginUrl
	return params
}
func (params *CreatePluginVersionParams) SetChangelog(Changelog string) *CreatePluginVersionParams {
	params.Changelog = &Changelog
	return params
}
func (params *CreatePluginVersionParams) SetPrivate(Private bool) *CreatePluginVersionParams {
	params.Private = &Private
	return params
}
func (params *CreatePluginVersionParams) SetCliVersion(CliVersion string) *CreatePluginVersionParams {
	params.CliVersion = &CliVersion
	return params
}
func (params *CreatePluginVersionParams) SetValidateStatus(ValidateStatus string) *CreatePluginVersionParams {
	params.ValidateStatus = &ValidateStatus
	return params
}

//
func (c *ApiService) CreatePluginVersion(PluginSid string, params *CreatePluginVersionParams) (*FlexV1PluginVersion, error) {
	path := "/v1/PluginService/Plugins/{PluginSid}/Versions"
	path = strings.Replace(path, "{"+"PluginSid"+"}", PluginSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Version != nil {
		data.Set("Version", *params.Version)
	}
	if params != nil && params.PluginUrl != nil {
		data.Set("PluginUrl", *params.PluginUrl)
	}
	if params != nil && params.Changelog != nil {
		data.Set("Changelog", *params.Changelog)
	}
	if params != nil && params.Private != nil {
		data.Set("Private", fmt.Sprint(*params.Private))
	}
	if params != nil && params.CliVersion != nil {
		data.Set("CliVersion", *params.CliVersion)
	}
	if params != nil && params.ValidateStatus != nil {
		data.Set("ValidateStatus", *params.ValidateStatus)
	}

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1PluginVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'FetchPluginVersion'
type FetchPluginVersionParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
}

func (params *FetchPluginVersionParams) SetFlexMetadata(FlexMetadata string) *FetchPluginVersionParams {
	params.FlexMetadata = &FlexMetadata
	return params
}

//
func (c *ApiService) FetchPluginVersion(PluginSid string, Sid string, params *FetchPluginVersionParams) (*FlexV1PluginVersion, error) {
	path := "/v1/PluginService/Plugins/{PluginSid}/Versions/{Sid}"
	path = strings.Replace(path, "{"+"PluginSid"+"}", PluginSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1PluginVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPluginVersion'
type ListPluginVersionParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPluginVersionParams) SetFlexMetadata(FlexMetadata string) *ListPluginVersionParams {
	params.FlexMetadata = &FlexMetadata
	return params
}
func (params *ListPluginVersionParams) SetPageSize(PageSize int) *ListPluginVersionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPluginVersionParams) SetLimit(Limit int) *ListPluginVersionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of PluginVersion records from the API. Request is executed immediately.
func (c *ApiService) PagePluginVersion(PluginSid string, params *ListPluginVersionParams, pageToken, pageNumber string) (*ListPluginVersionResponse, error) {
	path := "/v1/PluginService/Plugins/{PluginSid}/Versions"

	path = strings.Replace(path, "{"+"PluginSid"+"}", PluginSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
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

	ps := &ListPluginVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists PluginVersion records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPluginVersion(PluginSid string, params *ListPluginVersionParams) ([]FlexV1PluginVersion, error) {
	response, errors := c.StreamPluginVersion(PluginSid, params)

	records := make([]FlexV1PluginVersion, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams PluginVersion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPluginVersion(PluginSid string, params *ListPluginVersionParams) (chan FlexV1PluginVersion, chan error) {
	if params == nil {
		params = &ListPluginVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1PluginVersion, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePluginVersion(PluginSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPluginVersion(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamPluginVersion(response *ListPluginVersionResponse, params *ListPluginVersionParams, recordChannel chan FlexV1PluginVersion, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.PluginVersions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPluginVersionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPluginVersionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPluginVersionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPluginVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

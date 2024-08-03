/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateSafelist'
type CreateSafelistParams struct {
	// The phone number to be added in SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
}

func (params *CreateSafelistParams) SetPhoneNumber(PhoneNumber string) *CreateSafelistParams {
	params.PhoneNumber = &PhoneNumber
	return params
}

// Add a new phone number to SafeList.
func (c *ApiService) CreateSafelist(params *CreateSafelistParams) (*VerifyV2Safelist, error) {
	path := "/v2/SafeList/Numbers"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Safelist{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a phone number from SafeList.
func (c *ApiService) DeleteSafelist(PhoneNumber string) error {
	path := "/v2/SafeList/Numbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Check if a phone number exists in SafeList.
func (c *ApiService) FetchSafelist(PhoneNumber string) (*VerifyV2Safelist, error) {
	path := "/v2/SafeList/Numbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Safelist{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
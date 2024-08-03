/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// SupersimV1UsageRecord struct for SupersimV1UsageRecord
type SupersimV1UsageRecord struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that incurred the usage.
	AccountSid *string `json:"account_sid,omitempty"`
	// SID of a Sim resource to which the UsageRecord belongs. Value will only be present when either a value for the `Sim` query parameter is provided or when UsageRecords are grouped by `sim`. Otherwise, the value will be `null`.
	SimSid *string `json:"sim_sid,omitempty"`
	// SID of the Network resource the usage occurred on. Value will only be present when either a value for the `Network` query parameter is provided or when UsageRecords are grouped by `network`. Otherwise, the value will be `null`.
	NetworkSid *string `json:"network_sid,omitempty"`
	// SID of the Fleet resource the usage occurred on. Value will only be present when either a value for the `Fleet` query parameter is provided or when UsageRecords are grouped by `fleet`. Otherwise, the value will be `null`.
	FleetSid *string `json:"fleet_sid,omitempty"`
	// Alpha-2 ISO Country Code that the usage occurred in. Value will only be present when either a value for the `IsoCountry` query parameter is provided or when UsageRecords are grouped by `isoCountry`. Otherwise, the value will be `null`.
	IsoCountry *string `json:"iso_country,omitempty"`
	// The time period for which the usage is reported. The period is represented as a pair of `start_time` and `end_time` timestamps specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	Period *interface{} `json:"period,omitempty"`
	// Total data uploaded in bytes, aggregated by the query parameters.
	DataUpload *int64 `json:"data_upload,omitempty"`
	// Total data downloaded in bytes, aggregated by the query parameters.
	DataDownload *int64 `json:"data_download,omitempty"`
	// Total of data_upload and data_download.
	DataTotal *int64 `json:"data_total,omitempty"`
	// Total amount in the `billed_unit` that was charged for the data uploaded or downloaded. Will return 0 for usage prior to February 1, 2022. Value may be 0 despite `data_total` being greater than 0 if the data usage is still being processed by Twilio's billing system. Refer to [Data Usage Processing](https://www.twilio.com/docs/iot/supersim/api/usage-record-resource#data-usage-processing) for more details.
	DataTotalBilled *float32 `json:"data_total_billed,omitempty"`
	// The currency in which the billed amounts are measured, specified in the 3 letter ISO 4127 format (e.g. `USD`, `EUR`, `JPY`). This can be null when data_toal_billed is 0 and we do not yet have billing information for the corresponding data usage. Refer to [Data Usage Processing](https://www.twilio.com/docs/iot/supersim/api/usage-record-resource#data-usage-processing) for more details.
	BilledUnit *string `json:"billed_unit,omitempty"`
}

func (response *SupersimV1UsageRecord) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid      *string      `json:"account_sid"`
		SimSid          *string      `json:"sim_sid"`
		NetworkSid      *string      `json:"network_sid"`
		FleetSid        *string      `json:"fleet_sid"`
		IsoCountry      *string      `json:"iso_country"`
		Period          *interface{} `json:"period"`
		DataUpload      *int64       `json:"data_upload"`
		DataDownload    *int64       `json:"data_download"`
		DataTotal       *int64       `json:"data_total"`
		DataTotalBilled *interface{} `json:"data_total_billed"`
		BilledUnit      *string      `json:"billed_unit"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = SupersimV1UsageRecord{
		AccountSid:   raw.AccountSid,
		SimSid:       raw.SimSid,
		NetworkSid:   raw.NetworkSid,
		FleetSid:     raw.FleetSid,
		IsoCountry:   raw.IsoCountry,
		Period:       raw.Period,
		DataUpload:   raw.DataUpload,
		DataDownload: raw.DataDownload,
		DataTotal:    raw.DataTotal,
		BilledUnit:   raw.BilledUnit,
	}

	responseDataTotalBilled, err := client.UnmarshalFloat32(raw.DataTotalBilled)
	if err != nil {
		return err
	}
	response.DataTotalBilled = responseDataTotalBilled

	return
}

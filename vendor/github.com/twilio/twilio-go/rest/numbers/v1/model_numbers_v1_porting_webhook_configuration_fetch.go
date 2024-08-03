/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// NumbersV1PortingWebhookConfigurationFetch struct for NumbersV1PortingWebhookConfigurationFetch
type NumbersV1PortingWebhookConfigurationFetch struct {
	// The URL of the webhook configuration request
	Url *string `json:"url,omitempty"`
	// Webhook URL to send a request when a port in request or port in phone number event happens
	PortInTargetUrl *string `json:"port_in_target_url,omitempty"`
	// Webhook URL to send a request when a port out phone number event happens
	PortOutTargetUrl *string `json:"port_out_target_url,omitempty"`
	// List of notification events to send a request to the webhook URL
	NotificationsOf *[]string `json:"notifications_of,omitempty"`
	// Creation date for the port in webhook configuration
	PortInTargetDateCreated *time.Time `json:"port_in_target_date_created,omitempty"`
	// Creation date for the port out webhook configuration
	PortOutTargetDateCreated *time.Time `json:"port_out_target_date_created,omitempty"`
}
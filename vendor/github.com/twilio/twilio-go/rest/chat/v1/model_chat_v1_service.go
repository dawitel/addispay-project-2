/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
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

// ChatV1Service struct for ChatV1Service
type ChatV1Service struct {
	// The unique string that we created to identify the Service resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Service resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
	DefaultServiceRoleSid *string `json:"default_service_role_sid,omitempty"`
	// The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
	DefaultChannelRoleSid *string `json:"default_channel_role_sid,omitempty"`
	// The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
	DefaultChannelCreatorRoleSid *string `json:"default_channel_creator_role_sid,omitempty"`
	// Whether the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature is enabled. The default is `true`.
	ReadStatusEnabled *bool `json:"read_status_enabled,omitempty"`
	// Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Service instance. The default is `false`.
	ReachabilityEnabled *bool `json:"reachability_enabled,omitempty"`
	// How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds.
	TypingIndicatorTimeout *int `json:"typing_indicator_timeout,omitempty"`
	// DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
	ConsumptionReportInterval *int `json:"consumption_report_interval,omitempty"`
	// An object that describes the limits of the service instance. The `limits` object contains  `channel_members` to describe the members/channel limit and `user_channels` to describe the channels/user limit. `channel_members` can be 1,000 or less, with a default of 250. `user_channels` can be 1,000 or less, with a default value of 100.
	Limits *interface{} `json:"limits,omitempty"`
	// An object that contains information about the webhooks configured for this service.
	Webhooks *interface{} `json:"webhooks,omitempty"`
	// The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
	PreWebhookUrl *string `json:"pre_webhook_url,omitempty"`
	// The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
	PostWebhookUrl *string `json:"post_webhook_url,omitempty"`
	// The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookMethod *string `json:"webhook_method,omitempty"`
	// The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookFilters *[]string `json:"webhook_filters,omitempty"`
	// The notification configuration for the Service instance. See [Push Notification Configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more information.
	Notifications *interface{} `json:"notifications,omitempty"`
	// The absolute URL of the Service resource.
	Url *string `json:"url,omitempty"`
	// The absolute URLs of the Service's [Channels](https://www.twilio.com/docs/chat/api/channels), [Roles](https://www.twilio.com/docs/chat/api/roles), and [Users](https://www.twilio.com/docs/chat/api/users).
	Links *map[string]interface{} `json:"links,omitempty"`
}

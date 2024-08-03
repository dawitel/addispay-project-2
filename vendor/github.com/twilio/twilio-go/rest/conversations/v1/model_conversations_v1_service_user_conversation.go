/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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

// ConversationsV1ServiceUserConversation struct for ConversationsV1ServiceUserConversation
type ConversationsV1ServiceUserConversation struct {
	// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this User Conversation.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The number of unread Messages in the Conversation for the Participant.
	UnreadMessagesCount *int `json:"unread_messages_count,omitempty"`
	// The index of the last Message in the Conversation that the Participant has read.
	LastReadMessageIndex *int `json:"last_read_message_index,omitempty"`
	// The unique ID of the [participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) the user conversation belongs to.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The unique string that identifies the [User resource](https://www.twilio.com/docs/conversations/api/user-resource).
	UserSid *string `json:"user_sid,omitempty"`
	// The human-readable name of this conversation, limited to 256 characters. Optional.
	FriendlyName      *string `json:"friendly_name,omitempty"`
	ConversationState *string `json:"conversation_state,omitempty"`
	// Timer date values representing state update for this conversation.
	Timers *interface{} `json:"timers,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	Attributes *string `json:"attributes,omitempty"`
	// The date that this conversation was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this conversation was last updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Identity of the creator of this Conversation.
	CreatedBy         *string `json:"created_by,omitempty"`
	NotificationLevel *string `json:"notification_level,omitempty"`
	// An application-defined string that uniquely identifies the Conversation resource. It can be used to address the resource in place of the resource's `conversation_sid` in the URL.
	UniqueName *string `json:"unique_name,omitempty"`
	Url        *string `json:"url,omitempty"`
	// Contains absolute URLs to access the [participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) and [conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) of this conversation.
	Links *map[string]interface{} `json:"links,omitempty"`
}
package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"
	"github.com/dawitel/addispay-project-2/internal/utils"

	"github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var logger = utils.GetLogger()

// sendSMSNotifications sends sms notifications to the user
func sendSMSNotifications(notification *models.Notification) {
    accountSid := utils.GoDotEnvVariable("accountSid")      
    authToken := utils.GoDotEnvVariable("authToken")        
    fromPhoneNumber := notification.FromPhoneNumber        
    toPhoneNumber := notification.Merchant.MerchantPhone       
    messageBody := notification.MessageBody

    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })

    params := &openapi.CreateMessageParams{}
    params.SetTo(toPhoneNumber)
    params.SetFrom(fromPhoneNumber)
    params.SetBody(messageBody)

    resp, err := client.Api.CreateMessage(params)
    if err != nil {
        logger.Error("Failed to send SMS notification to: ", toPhoneNumber, err)
    }

    logger.Printf("SMS Message sent, SID: %s\n", *resp.Sid)
}



func sendEmailNotification(notification *models.Notification) {
	
    from := mail.NewEmail(notification.SenderName, notification.SenderEmail)
    to := mail.NewEmail("Recipient", notification.RecipientEmail)

    message := mail.NewSingleEmail(from, notification.Subject, to, notification.PlainTextContent, notification.HtmlContent)
    API_KEY := utils.GoDotEnvVariable("API_KEY_MAILER")
    client := sendgrid.NewSendClient(API_KEY)

    response, err := client.Send(message)
    if err != nil {
        logger.Error("Failed to send notification email to ", notification.RecipientEmail, err)
    } else {
       logger.Success("Email sent successfully with status code: ", response.StatusCode)
    }
}

// sendNotificationsToAPIgateway sends the notifications to the frontend app via the API gateway
func sendNotificationsToAPIgateway(notification *models.Notification) {
	// Load configuration files to the environment
	config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files")
    }
	// marshal the data to JSON
	notificationJSON, err := json.Marshal(notification)
	if err!= nil {
		logger.Error("Failed to marshal the notificaion messages: ", err)
	}

	URL := fmt.Sprintf("%s/api/v1/notifications", config.APIGatewayAddr)
	
	// send the JSON data to the API gateway
	response, err := http.Post(URL, "application/json", bytes.NewBuffer(notificationJSON))
	if err != nil {
		logger.Error("Failed to send the order Response to the API gateway: ", err)
	}

	defer response.Body.Close()
	logger.Success("Order Response data sent to the API gateway: ", response.StatusCode)
} 
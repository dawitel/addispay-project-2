package models

type Notification struct {
	RecipientEmail 		string
	SenderName 			string "Dawit Elias"
	SenderEmail 		string "dawiteliaskassaye@gmail.com"
	Subject 			string "🎆You have a New successful purchase"
	PlainTextContent 	string "The order you recieved from orderID={] has been completed successfully. You are crushing it this week!"
	HtmlContent 		string "<strong>The order you recieved from orderID={] has been completed successfully. You are crushing it this week!</strong>"
	FromPhoneNumber 	string "+19789938396"
	ToPhoneNumber 		string 
	MessageBody  		string
	Merchant            Merchant
}

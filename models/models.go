package models

// Sample struct for email
type EmailStruct struct {
	ReceiverEmail string `json:"email"`
	EmailSubject  string `json:"emailSubject"`
	EmailBody     string `json:"emailBody"`
}

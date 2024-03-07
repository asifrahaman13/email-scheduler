package models

type EmailStruct struct {
	ReceiverEmail string `json:"email"`
	EmailSubject  string `json:"emailSubject"`
	EmailBody     string `json:"emailBody"`
}

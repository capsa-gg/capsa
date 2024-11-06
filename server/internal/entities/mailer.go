package entities

// EmailSender is used to make email implementations hidden in calling code.
type EmailSender interface {
	// SendEmail sends an email with the provided arguments.
	SendEmail(toEmail, toName, subject, htmlContents string) error
}

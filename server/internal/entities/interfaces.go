package entities

// EmailSender is used to make email implementations hidden in calling code.
// Note: this must be placed in entities, not in interactors, as that would cause a circular import.
type EmailSender interface {
	// SendEmail sends an email with the provided arguments.
	SendEmail(toEmail, toName, subject, htmlContents string) error
}

package emails

import (
	"fmt"

	"go.uber.org/zap"
)

func returnHTMLGenerationError(log *zap.SugaredLogger, err error) error {
	log.Errorf("html generation error: %s", err)

	return fmt.Errorf("error generating email html: %w", err)
}

func returnEmailSendingError(log *zap.SugaredLogger, err error) error {
	log.Errorf("email sending error: %s", err)

	return fmt.Errorf("error sending email: %w", err)
}

// SendAccountSetPassword sends an email to a user, prompting them to set a password for a newly created account.
func (e Emails) SendAccountSetPassword(toEmail, firstName, token string) error {
	log := e.logger.Named("SendAccountSetPassword").With("to_email", toEmail)

	html, err := e.hermes.GenerateHTML(setPasswordTemplate(firstName, token, e.generatePasswordResetLink(token)))
	if err != nil {
		return returnHTMLGenerationError(log, err)
	}

	err = e.mailer.SendEmail(toEmail, firstName, "Please set your Capsa password", html)
	if err != nil {
		return returnEmailSendingError(log, err)
	}

	return nil
}

// SendLoginSuccessNotification sends an email notifying a user of a successful login.
func (e Emails) SendLoginSuccessNotification(toEmail, firstName string) error {
	log := e.logger.Named("SendLoginSuccessNotification").With("to_email", toEmail)

	html, err := e.hermes.GenerateHTML(userLoginSuccessTemplate(firstName))
	if err != nil {
		return returnHTMLGenerationError(log, err)
	}

	err = e.mailer.SendEmail(toEmail, firstName, "Successful account login", html)
	if err != nil {
		return returnEmailSendingError(log, err)
	}

	return nil
}

// SendPasswordResetToken sends an email with a password reset code.
func (e Emails) SendPasswordResetToken(toEmail, firstName, token string) error {
	log := e.logger.Named("SendPasswordResetToken").With("to_email", toEmail)

	html, err := e.hermes.GenerateHTML(resetPasswordCodeTemplate(firstName, token, e.generatePasswordResetLink(token)))
	if err != nil {
		return returnHTMLGenerationError(log, err)
	}

	err = e.mailer.SendEmail(toEmail, firstName, "Your Capsa password reset code", html)
	if err != nil {
		return returnEmailSendingError(log, err)
	}

	return nil
}

// SendPasswordResetConfirmation sends an email confirming the password has been reset.
func (e Emails) SendPasswordResetConfirmation(toEmail, firstName string) error {
	log := e.logger.Named("SendPasswordResetToken").With("to_email", toEmail)

	html, err := e.hermes.GenerateHTML(resetPasswordConfirmationTemplate(firstName))
	if err != nil {
		return returnHTMLGenerationError(log, err)
	}

	err = e.mailer.SendEmail(toEmail, firstName, "Your Capsa password has been reset", html)
	if err != nil {
		return returnEmailSendingError(log, err)
	}

	return nil
}

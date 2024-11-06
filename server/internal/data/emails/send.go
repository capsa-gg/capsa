package emails

import "fmt"

// SendAccountSetPassword sends an email to a user, prompting them to set a password for a newly created account.
func (e Emails) SendAccountSetPassword(toEmail, firstName, token string) error {
	log := e.logger.Named("SendAccountSetPassword").With("to_email", toEmail)

	html, err := e.hermes.GenerateHTML(setPasswordTemplate(firstName, token))
	if err != nil {
		log.Errorf("html generation error: %s", err)

		return fmt.Errorf("error generating email html: %w", err)
	}

	err = e.mailer.SendEmail(toEmail, firstName, "Please set your Capsa password", html)

	if err != nil {
		log.Errorf("email sending error: %s", err)

		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

// SendLoginSuccessNotification sends an email notifying a user of a successful login.
func (e Emails) SendLoginSuccessNotification(toEmail, firstName string) error {
	log := e.logger.Named("SendLoginSuccessNotification").With("to_email", toEmail)

	html, err := e.hermes.GenerateHTML(userLoginSuccessTemplate(firstName))
	if err != nil {
		log.Errorf("html generation error: %s", err)

		return fmt.Errorf("error generating email html: %w", err)
	}

	err = e.mailer.SendEmail(toEmail, firstName, "Successful account login", html)

	if err != nil {
		log.Errorf("email sending error: %s", err)

		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

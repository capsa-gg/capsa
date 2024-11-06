package mailer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const brevoAPIEndpoint = "https://api.brevo.com/v3/smtp/email"

type contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type emailRequest struct {
	Sender      contact   `json:"sender"`
	To          []contact `json:"to"`
	Subject     string    `json:"subject"`
	HTMLContent string    `json:"htmlContent"`
}

// SendEmail sends an email with the provided arguments.
func (m Mailer) SendEmail(toEmail, toName, subject, htmlContents string) error {
	log := m.logger.Named("SendEmail").Sugar().With(
		"to_email", toEmail,
		"to_name", toName,
		"subject", subject)

	log.Debug("start sending email")

	reqBody := emailRequest{
		Sender: contact{
			Name:  m.config.EmailSenderName,
			Email: m.config.EmailSenderEmail,
		},
		To:          []contact{{Name: toName, Email: toEmail}},
		Subject:     subject,
		HTMLContent: htmlContents,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error converting email request body to json: %w", err)
	}

	payload := strings.NewReader(string(jsonBody))

	req, err := http.NewRequest("POST", brevoAPIEndpoint, payload)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", m.config.BrevoAPIKey)

	log.Debug("start sending email")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing email sending request: %w", err)
	}

	log.Debug("email request made")

	defer func() {
		if errDefer := res.Body.Close(); errDefer != nil {
			log.Errorf("error when closing res.Body: %s", errDefer)
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading email sending response body: %w", err)
	}

	log.Debugf("received email request response: %s", body)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("did not receive 2xx status code while sending email, got %d", res.StatusCode)
	}

	log.Info("email sent successfully")

	return nil
}

package emails

import "github.com/matcornic/hermes/v2"

func addDefaultValuesToTemplate(template *hermes.Email) {
	template.Body.Greeting = "Dear"
	template.Body.Signature = "With kind regards"
}

func setPasswordTemplate(firstName, resetToken string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Your Capsa account has been created.",
				"Please use the code " + resetToken + " to set your password.",
			},
			// TODO: button element
			Outros: []string{
				"If you don't set your password before the code expires, you can request a password reset to set your password.",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

func userLoginSuccessTemplate(firstName string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"There has been a successful login attempt with your Capsa account.",
				"In case this was not you, please reach out to your Capsa admin as soon as possible.",
			},
			Outros: []string{
				"Do you need help, or do you have a question? Feel free to reach out by replying to this email, we are happy to help out!",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

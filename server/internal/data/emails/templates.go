package emails

import "github.com/matcornic/hermes/v2"

func addDefaultValuesToTemplate(template *hermes.Email) {
	template.Body.Greeting = "Dear"
	template.Body.Signature = "With kind regards"
}

func setPasswordTemplate(firstName, resetToken, resetLink string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Your Capsa account has been created.",
				"The code is " + resetToken + " to set your password.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please set your password here:",
					Button: hermes.Button{
						Text: "Set your password",
						Link: resetLink,
					},
				},
			},
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

func resetPasswordCodeTemplate(firstName, resetToken, resetLink string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Reset your Capsa password.",
				"Please use the code " + resetToken + " to reset your password.",
				"If you have not requested your password reset, you can ignore this email.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Set your new password here:",
					Button: hermes.Button{
						Text: "Reset password",
						Link: resetLink,
					},
				},
			},
			Outros: []string{
				"If you don't reset your password before the code expires, you need to request a new code.",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

func resetPasswordConfirmationTemplate(firstName string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Your Capsa password has been reset.",
				"If you have not performed this action, please reach out to your Capsa administrator right away.",
			},
			Outros: []string{
				"Do you need help, or do you have a question? Feel free to reach out by replying to this email, we are happy to help out!",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

package i18n

func EnUS() map[string]string {
	return map[string]string{
		"400_bad_request":              "The request cannot be performed because of malformed or missing parameters.",
		"401_unauthorized":             "Unauthorized. Please Re-Login",
		"403_forbidden":                "The user does not have permission to :action.",
		"404_not_found":                "The resource you have specified cannot be found.",
		"500_internal_error":           "Failed to connect to the server, please try again later.",
		"invalid_username_or_password": "Invalid username or password",
	}
}

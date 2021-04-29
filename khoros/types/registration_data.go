package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/registration_data
// RegistrationData refers to the properties from the returned `registration_data` object
type RegistrationData struct {
	ConfirmEmailStatus      bool   `json:"confirm_email_status"`
	RegistrationAccessLevel string `json:"registration_access_Level"` // string enum
	RegistrationTime        string `json:"registration_time"`         // JSON DateTime
	Status                  string `json:"status"`                    // string enum
}

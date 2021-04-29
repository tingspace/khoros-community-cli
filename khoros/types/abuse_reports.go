package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/abuse_reports
// AbuseReports refers to the properties from the returned `abuse_reports` object
type AbuseReports struct {
	Body     string    `json:"body"`
	Image    *Image    `json:"image"`
	Message  *Messages `json:"message"`
	Reporter *Users    `json:"reporter"`
}

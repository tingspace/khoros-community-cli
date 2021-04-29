package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/me_toos
// MeToos refers to the properties from the returned `me_toos` object
type MeToos struct {
	Href    string    `json:"href"`
	Id      string    `json:"id"`
	Message *Messages `json:"message"`
	Time    string    `json:"time"` // JSON DateTime
	User    *Users    `json:"user"`
}

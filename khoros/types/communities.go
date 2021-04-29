package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/communities
// Communities refers to the properties from the returned `communities` object
type Communities struct {
	Ancestors           []*Nodes    `json:"ancestors"`
	Announcements       string      `json:"announcements"`
	DatePattern         string      `json:"date_pattern"`
	Description         string      `json:"description"`
	FeatureContext      interface{} `json:"feature_context"` // Object not documented
	FriendlyDateEnabled bool        `json:"friendly_date_enabled"`
	FriendlyDateMaxAge  int         `json:"friendly_date_max_age"`
	Href                string      `json:"href"`
	Id                  string      `json:"id"`
	Language            string      `json:"language"`
	MobileAnnouncements string      `json:"mobile_announcements"`
	NodeType            string      `json:"node_type"`
	ShortTitle          string      `json:"short_title"`
	Title               string      `json:"title"`
	UserContext         string      `json:"user_context"`
	ViewHref            string      `json:"view_href"`
	WebUI               interface{} `json:"webui"` // Object not documented
}

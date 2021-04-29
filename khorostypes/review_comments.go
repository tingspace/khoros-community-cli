package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/review_comments
// ReviewComments refers to the properties from the returned `review_comments` object
type ReviewComments struct {
	Author           *Users      `json:"author"`
	Body             string      `json:"body"`
	CurrentRevision  *Revisions  `json:"current_revision"`
	Href             string      `json:"href"`
	Id               string      `json:"id"`
	Language         string      `json:"language"`
	LastPublishTime  string      `json:"last_publish_time"`
	Metrics          interface{} `json:"metrics"`            // Object undocumented
	ModerationStatus interface{} `json:"moderation_status"`  // Object undocumented
	Placeholder      interface{} `json:"placeholder"`        // Object undocumented
	PostTime         string      `json:"post_time"`          // JSON DateTime
	PostTimeFriendly string      `json:"post_time_friendly"` // JSON DateTime
	Product          *Products   `json:"product"`
	ReadOnly         bool        `json:"read_only"`
	Review           *Reviews    `json:"review"`
	ViewHref         string      `json:"view_href"`
}

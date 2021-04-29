package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/reviews
// Reviews refers to the properties from the returned `reviews` object
type Reviews struct {
	Author           *Users                  `json:"author"`
	Body             string                  `json:"body"`
	Comments         []*ReviewComments       `json:"comments"`
	CurrentRevision  *Revisions              `json:"current_revision"`
	Href             string                  `json:"href"`
	Id               string                  `json:"id"`
	Images           string                  `json:"images"`
	Language         string                  `json:"language"`
	LastPublishTime  string                  `json:"last_publish_time"` // JSON DateTime
	Metrics          interface{}             `json:"metrics"`
	ModerationStatus string                  `json:"moderation_status"`  // Hopefully a strin enum
	Placeholder      interface{}             `json:"placeholder"`        // Object not documented
	PostTime         string                  `json:"post_time"`          // JSON DateTime
	PostTimeFriendly string                  `json:"post_time_friendly"` // JSON DateTime
	Product          *Products               `json:"product"`
	ReadOnly         bool                    `json:"read_only"`
	ReviewDimensions *ReviewDimensionRatings `json:"review_dimensions"`
	ReviewRating     *ReviewRatings          `json:"review_rating"`
	Subject          string                  `json:"subject"`
	Videos           []*Videos               `json:"videos"`
	ViewHref         string                  `json:"view_href"`
}

package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/tkb_helpfulness_ratings
// TKBHelpfulnessRatings refers to the properties from the returned `tkb_helpfulness_ratings` object
type TKBHelpfulnessRatings struct {
	Id      string    `json:"id"`
	Message *Messages `json:"message"`
	Time    string    `json:"time"` // JSON DateTime
	User    *Users    `json:"user"`
	Value   bool      `json:"value"`
}

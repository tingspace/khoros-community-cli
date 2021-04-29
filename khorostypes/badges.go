package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/badges
// Badges refers to the properties from the returned `badges` object
type Badges struct {
	ActivationDate string `json:"activation_date"` // JSON DateTime
	Awarded        int    `json:"awarded"`
	Description    string `json:"description"`
	IconUrl        string `json:"icon_url"`
	Id             string `json:"id"`
	Title          string `json:"title"`
}

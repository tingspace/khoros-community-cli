package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/review_dimension_ratings
// Revisions refers to the properties from the returned `revisions` object
type Revisions struct {
	Id             string `json:"id"`
	LastEditAuthor *Users `json:"last_edit_author"`
	LastEditTime   string `json:"last_edit_time"` // JSON DateTime
}

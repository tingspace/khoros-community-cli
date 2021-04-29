package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/review_ratings
// ReviewRatings refers to the properties from the returned `review_ratings` object
type ReviewRatings struct {
	Id      string    `json:"id"`
	Product *Products `json:"product"`
	Review  *Reviews  `json:"review"`
	Time    string    `json:"time"` // JSON DateTime
	Value   int       `json:"value"`
}

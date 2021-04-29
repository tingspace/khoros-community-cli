package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/ratings
// Ratings refers to the properties from the returned `ratings` object
type Ratings struct {
	Id      string    `json:"id"`
	Message *Messages `json:"message"`
	Time    string    `json:"time"`
	User    *Users    `json:"user"`
	Value   int       `json:"value"`
}

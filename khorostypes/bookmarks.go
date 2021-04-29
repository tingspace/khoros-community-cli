package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/bookmarks
// Bookmarks refers to the properties from the returned `bookmarks` object
type Bookmarks struct {
	Href       string      `json:"href"`
	Id         string      `json:"id"`
	Subscriber *Users      `json:"subscriber"`
	Target     interface{} `json:"target"` // Object not documented
}

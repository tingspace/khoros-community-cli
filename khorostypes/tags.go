package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/tags
// Tags refers to the properties from the returned `tags` object
type Tags struct {
	Href     string      `json:"href"`
	Id       string      `json:"id"`
	Messages []*Messages `json:"messages"`
	Text     string      `json:"text"`
	Time     string      `json:"time"` // JSON DateTime
}

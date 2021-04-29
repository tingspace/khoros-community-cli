package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/labels
// Labels refers to the properties from the returned `labels` object
type Labels struct {
	Href     string      `json:"href"`
	Id       string      `json:"id"`
	Messages []*Messages `json:"messages"`
	Text     string      `json:"text"`
	Time     string      `json:"time"` // JSON DateTime
	User     *Users      `json:"user"`
}

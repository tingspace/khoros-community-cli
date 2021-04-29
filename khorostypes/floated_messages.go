package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/floated_messages
// FloatMessages refers to the properties from the returned `floated_messages` object
type FloatedMessages struct {
	Href    string      `json:"href"`
	Id      string      `json:"id"`
	Message *Messages   `json:"message"`
	Scope   interface{} `json:"scope"` // Object not documented
	User    *Users      `json:"user"`
}

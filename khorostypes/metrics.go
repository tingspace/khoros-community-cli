package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/metrics
// Metrics refers to the properties from the returned `metrics` object
type Metrics struct {
	Board    *Boards     `json:"board"`
	Category *Categories `json:"category"`
	Href     string      `json:"href"`
	Id       string      `json:"id"`
	Role     *Roles      `json:"role"`
	User     *Users      `json:"user"`
	Value    int         `json:"value"`
}

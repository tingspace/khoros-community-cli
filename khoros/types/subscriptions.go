package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/subscriptions
// Subscriptions refers to the properties from the returned `subscriptions` object
type Subscriptions struct {
	Href       string      `json:"href"`
	Id         string      `json:"id"`
	Node       *Nodes      `json:"node"`
	Subscriber *Users      `json:"subscriber"`
	Target     interface{} `json:"target"` // Object not documented
}

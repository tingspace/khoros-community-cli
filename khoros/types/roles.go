package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/roles
// Roles refers to the properties from the returned `roles` object
type Roles struct {
	Board      *Boards     `json:"board"`
	Category   *Categories `json:"category"`
	Href       string      `json:"href"`
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Node       *Nodes      `json:"node"`
	NodeType   string      `json:"node_type"`
	RoleStatus string      `json:"role_status"` // Hoping this is a string enum
	Users      []*Users    `json:"users"`
}

package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/grouphubs
// GroupHubs refers to the properties from the returned `group_hubs` object
type GroupHubs struct {
	Ancestors       []*Nodes    `json:"ancestors"`
	Avatar          interface{} `json:"avatar"` // Object not documented
	Children        []*Nodes    `json:"children"`
	CreationDate    string      `json:"creation_date"` // JSON DateTime
	Depth           int         `json:"depth"`
	Descendants     []*Nodes    `json:"descendants"`
	Description     string      `json:"description"`
	Href            string      `json:"href"`
	Id              string      `json:"id"`
	MembershipType  string      `json:"membership_type"`  // string enum
	MessageActivity interface{} `json:"message_activity"` // Object not documented
	Messages        []*Messages `json:"messages"`
	NodeType        string      `json:"node_type"` // string enum
	Parent          *Nodes      `json:"parent"`
	ParentCategory  *Categories `json:"parent_category"`
	Position        int         `json:"position"`
	RootCategory    *Categories `json:"root_category"`
	ShortTitle      string      `json:"short_title"`
	Title           string      `json:"title"`
	Topics          []*Messages `json:"topics"`
	UserContext     interface{} `json:"user_context"` // Object not documented
	ViewHref        string      `json:"view_href"`
	Views           int         `json:"views"`
}

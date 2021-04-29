package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/nodes-1
// Nodes refers to the properties from the returned `notes` object
type Nodes struct {
	Ancestors            []*Nodes    `json:"ancestors"`
	Children             []*Nodes    `json:"children"`
	ConversationStyle    string      `json:"conversation_style"`
	CreationDate         string      `json:"creation_date"`          // JSON DateTime
	CreationDateFriendly string      `json:"creation_date_friendly"` // JSON DateTime
	Depth                int         `json:"depth"`
	Descendants          []*Nodes    `json:"descendants"`
	Description          string      `json:"description"`
	Hidden               bool        `json:"hidden"`
	Href                 string      `json:"href"`
	Id                   string      `json:"id"`
	JoinDate             string      `json:"join_date"` // JSON DateTime
	Members              []*Users    `json:"members"`
	MessageActivity      interface{} `json:"message_activity"` // Object not documneted >:(
	Messages             []*Messages `json:"messages"`
	NodeType             string      `json:"node_type"`
	Parent               *Nodes      `json:"parent"`
	Position             int         `json:"position"`
	Roles                *Roles      `json:"roles"`
	RootCategory         *Categories `json:"root_category"`
	ShortTitle           string      `json:"short_title"`
	Title                string      `json:"title"`
	Topics               []*Messages `json:"topics"`
	UserContext          interface{} `json:"user_context"` // Object not documented >:(
	ViewHref             string      `json:"view_href"`
	Views                int         `json:"views"`
}

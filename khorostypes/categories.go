package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/categories
// Categories refers to the properties from the returned `categories` object
type Categories struct {
	AncestorCategories   []*Categories `json:"ancestor_categories"`
	Announcements        string        `json:"announcements"`
	Boards               []*Boards     `json:"boards"`
	ChildCategories      []*Categories `json:"child_categories"`
	CreationDate         string        `json:"creation_date"`          // JSON DateTime
	CreationDateFriendly string        `json:"creation_date_friendly"` // JSON DateTime
	DatePattern          string        `json:"date_pattern"`
	Depth                int           `json:"depth"`
	DescendentCategories []*Categories `json:"descendant_categories"`
	Description          string        `json:"description"`
	FriendlyDateEnabled  bool          `json:"friendly_date_enabled"`
	FriendlyDateMaxAge   bool          `json:"friendly_date_max_age"`
	Hidden               bool          `json:"hidden"`
	Href                 string        `json:"href"`
	Id                   string        `json:"id"`
	Language             string        `json:"language"`
	MessageActivity      interface{}   `json:"message_activity"` // Object not documented >:(
	Messages             []*Messages   `json:"messages"`
	MobileAnnouncements  string        `json:"mobile_announcements"`
	NodeType             string        `json:"node_type"`
	ParentCategory       *Categories   `json:"parent_category"`
	Position             int           `json:"position"`
	RootCategory         *Categories   `json:"root_category"`
	ShortTitle           string        `json:"short_title"`
	Skin                 string        `json:"skin"`
	Title                string        `json:"title"`
	Topics               string        `json:"topics"`
	UserContext          interface{}   `json:"user_context"` // Object not documented >:(
	ViewHref             string        `json:"view_href"`
	Views                int           `json:"views"`
}

package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/tagged_product_scope
// TaggedProductScope refers to the properties from the returned `tagged_product_scope` object
type TaggedProductScope struct {
	Board        *Boards      `json:"board"`
	Category     *Categories  `json:"category"`
	Community    *Communities `json:"community"`
	Conversation interface{}  `json:"conversation"` // Object not documented
	GroupHub     *GroupHubs   `json:"group_hub"`
	Message      *Messages    `json:"message"`
	Occurs       int          `json:"occurs"`
}

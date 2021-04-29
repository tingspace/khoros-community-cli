package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/custom_tag_message_scope
// CustomTagMessageScope refers to the properties from the returned `custom_tag_message_scope` object
type CustomTagMessageScope struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/custom_tags
// CustomTags refers to the properties from the returned `custom_tags` object
type CustomTags struct {
	Href           string                 `json:"href"`
	Id             string                 `json:"id"`
	MessageScope   *CustomTagMessageScope `json:"message_scope"`
	Messages       []*Messages            `json:"messages"`
	PossibleValues string                 `json:"possible_values"`
	Text           string                 `json:"text"`
}

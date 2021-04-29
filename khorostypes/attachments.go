package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/attachments
// Attachments refers to the properties from the returned `attachments` object
type Attachments struct {
	ContentType string    `json:"content_type"`
	FileName    string    `json:"filename"`
	FileSize    int       `json:"filesize"`
	Href        string    `json:"href"`
	Id          string    `json:"id"`
	Message     *Messages `json:"message"`
	Position    int       `json:"position"`
	Url         string    `json:"url"`
}

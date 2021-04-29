package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/message_statuses
// MessageStatus refers to properties within the LiQL Object
type MessageStatus struct {
	Completed bool   `json:"completed"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	TypeKey   string `json:"type_key"`
}

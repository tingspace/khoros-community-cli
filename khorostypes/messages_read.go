package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/messages_read
// MessagesRead refers to the properties from the returned `messages_read` object
type MessagesRead struct {
	MarkUnread bool     `json:"mark_unread"`
	MessageId  string   `json:"message_id"`
	MessageIds []string `json:"message_ids"` // Hoping this is an array of strings
	TopicId    string   `json:"topic_id"`
	User       *Users   `json:"user"`
}

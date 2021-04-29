package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/outbox_notes
// OutboxNotes refers to the properties from the returned `outbox_notes` object
type OutboxNotes struct {
	Body         string        `json:"body"`
	Id           string        `json:"id"`
	IsRead       bool          `json:"is_read"`
	ReceivedTime string        `json:"received_time"` // JSON DateTime
	Recipients   []interface{} `json:"recipients"`    // Object not documented
	SentTime     string        `json:"sent_time"`     // JSON DateTime
	Subject      string        `json:"subject"`
	User         *Users        `json:"user"`
}

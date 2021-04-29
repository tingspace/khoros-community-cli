package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/inbox_notes
// InboxNotes refers to the properties from the returned `inbox_notes` object
type InboxNotes struct {
	Body         string `json:"body"`
	Id           string `json:"id"`
	IsRead       bool   `json:"is_read"`
	ReceivedTime string `json:"received_time"` // JSON DateTime
	Sender       *Users `json:"sender"`
	SentTime     string `json:"sent_time"` // JSON DateTime
	Subject      string `json:"subject"`
}

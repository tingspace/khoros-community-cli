package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/threaded_notes
// ThreadedNotes refers to the properties from the returned `threaded_notes` object
type ThreadedNotes struct {
	Body         string        `json:"body"`
	Id           string        `json:"id"`
	ReceivedTime string        `json:"received_time"` // JSON DateTime
	Recipients   []interface{} `json:"recipients"`    // Object not documented
	Sender       *Users        `json:"sender"`
	SentTime     string        `json:"sent_time"`
}

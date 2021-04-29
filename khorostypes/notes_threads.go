package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/notes_threads
// NotesThreads refers to the properties from the returned `notes_threads` object
type NotesThreads struct {
	Id            string         `json:"id"`
	LatestMessage *ThreadedNotes `json:"latest_message"`
	Subject       string         `json:"subject"`
	ThreadType    string         `json:"thread_type"`  // string enum
	UpdatedTime   string         `json:"updated_time"` // JSON DateTime
}

package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/occasion_data
// OccasionData refers to the properties from the returned `occasion_data` object
type OccasionData struct {
	StartTime      string      `json:"start_time"` // JSON DateTime
	EndTime        string      `json:"end_time"`   // JSON DateTime
	TimeZone       string      `json:"timezone"`
	Status         string      `json:"status"` // string enum
	Location       string      `json:"location"`
	IsLiveStream   bool        `json:"is_live_stream"`
	LiveStreamUrl  string      `json:"live_stream_url"`
	FeaturedGuests []*Users    `json:"featured_guests"`
	RSVP           interface{} `json:"rsvp"`            // Object not documented
	PendingInvites interface{} `json:"pending_invites"` // Object not documented
}

package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/rsvps
// RSVPs refers to the properties from the returned `rsvps` object
type RSVPs struct {
	Id           string `json:"id"`
	User         *Users `json:"user"`
	RSVPResponse string `json:"rsvp_response"` // string enum
}

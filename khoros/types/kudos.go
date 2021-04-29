package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/kudos
// Kudos refers to the Object property used by collections
type Kudos struct {
	Href    string    `json:"href"`
	Id      string    `json:"id"`
	Message *Messages `json:"message"`
	Time    string    `json:"time"` // JSON DateTime
	User    *Users    `json:"user"`
	Weight  int       `json:"weight"`
}

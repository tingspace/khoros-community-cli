package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/solutions_data
// SolutionsData refers to the properties from the returned `solutions_data` object
type SolutionsData struct {
	Accepter  *Users `json:"accepter"`
	MessageId string `json:"message_id"`
	Time      string `json:"time"` // JSON DateTime
}

package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/images
// Image refers to the Object property used by collections
type Image struct {
	Album            *Albums     `json:"album"`
	Comments         []*Messages `json:"comments"`
	Description      string      `json:"description"`
	Height           int         `json:"height"`
	Id               string      `json:"id"`
	Messages         []*Messages `json:"messages"`
	ModerationStatus string      `json:"moderation_status"`
	Owner            *Users      `json:"owner"`
	Title            string      `json:"title"`
	Visibility       string      `json:"visibility"`
	Width            int         `json:"width"`
}

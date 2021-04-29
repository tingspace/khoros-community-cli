package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/videos
// Videos refers to the properties from the returned `videos` object
type Videos struct {
	Description      string      `json:"description"`
	Duration         int         `json:"duration"`
	Format           string      `json:"format"`
	Height           int         `json:"height"`
	Href             string      `json:"href"`
	Id               string      `json:"id"`
	Messages         []*Messages `json:"messages"`
	ModerationStatus interface{} `json:"moderation_status"` // Object not documented
	OoyalaId         string      `json:"ooyala_id"`
	Owner            *Users      `json:"owner"`
	PrivacyLevel     interface{} `json:"privacy_level"` // Object not documented
	Size             int         `json:"size"`
	ThumbHref        string      `json:"thumb_href"`
	Title            string      `json:"title"`
	UploadDate       string      `json:"upload_date"` // JSON DateTime
	VideoType        string      `json:"video_type"`
	ViewHref         string      `json:"view_href"`
	Width            int         `json:"width"`
}

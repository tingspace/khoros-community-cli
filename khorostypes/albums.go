package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/albums
// Albums refers to the properties from the returned `albums` object
type Albums struct {
	Cover        *Image `json:"cover"`
	Default      bool   `json:"default"`
	Description  string `json:"description"`
	Href         string `json:"href"`
	Id           string `json:"id"`
	Images       *Image `json:"images"`
	Owner        *Users `json:"owner"`
	PrivacyLevel string `json:"privacy_level"`
	Title        string `json:"title"`
	ViewHref     string `json:"view_href"`
}

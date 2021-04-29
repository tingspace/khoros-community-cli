package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/boards
// Boards refers to the properties from the returned `board` object
type Boards struct {
	AllowedLabels          string           `json:"allowed_labels"`
	Announcements          string           `json:"announcements"`
	AvailableStatuses      []*MessageStatus `json:"available_statuses"`
	CommentsEnabled        bool             `json:"comments_enabled"`
	ConversationStyle      string           `json:"conversation_style"`
	CreationDate           string           `json:"creation_date"`          // JSON DateTime is a string
	CreationDateFriendly   string           `json:"creation_date_friendly"` // JSON DateTime is a string
	DatePattern            string           `json:"date_pattern"`
	Depth                  int              `json:"depth"`
	Description            string           `json:"description"`
	FriendlyDateEnabled    bool             `json:"friendly_date_enabled"`
	FriendlyDateMaxAge     int              `json:"friendly_date_max_age"`
	Hidden                 bool             `json:"hidden"`
	Href                   string           `json:"href"`
	Id                     string           `json:"id"`
	ImagePrivacy           string           `json:"image_privacy"`
	Language               string           `json:"language"`
	MediaType              string           `json:"media_type"`
	MessageActivity        interface{}      `json:"message_activity"` // Object not documented >:(
	Messages               []*Messages      `json:"messages"`
	MobileAnnouncements    string           `json:"mobile_announcements"`
	Nested                 bool             `json:"nested"`
	NodeType               string           `json:"node_type"`
	OneEntryPerContest     bool             `json:"one_entry_per_contest"`
	OneKudoPerContest      bool             `json:"one_kudo_per_contest"`
	ParentCategory         *Categories      `json:"parent_category"`
	Position               int              `json:"position"`
	PostingDateEnd         string           `json:"posting_date_end"`   // JSON DateTime
	PostingDateStart       string           `json:"posting_date_start"` // JSON DateTime
	Rating                 string           `json:"rating"`
	RequireThreadRootLabel string           `json:"require_thread_root_label"`
	RootCategory           *Categories      `json:"root_category"`
	ShortTitle             string           `json:"short_title"`
	Skin                   string           `json:"skin"`
	Title                  string           `json:"title"`
	Topics                 []*Messages      `json:"topics"`
	UserContext            interface{}      `json:"user_context"` // Object not documented >:(
	ViewHref               string           `json:"view_href"`
	VotingDateEnd          string           `json:"voting_date_end"`       // JSON DateTime
	VotingDateStart        string           `json:"voting_date_start"`     // JSON DateTime
	WinnerAnnouncedDate    string           `json:"winner_announced_date"` // JSON DateTime
}

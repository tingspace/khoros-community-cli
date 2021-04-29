package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/messages
// Messages refers to the properties from the returned `messages` object
type Messages struct {
	Attachments                   []*Attachments         `json:"attachments"`
	Author                        *Users                 `json:"author"`
	Board                         *Boards                `json:"board"`
	BoardRelativeId               string                 `json:"board_relative_id"`
	Body                          string                 `json:"body"`
	CanAcceptSolution             bool                   `json:"can_accept_solution"`
	CanonicalUrl                  string                 `json:"canonical_url"`
	ContentWorkflow               *ContentWorkflow       `json:"content_workflow"`
	ContextId                     string                 `json:"context_id"`
	ContextUrl                    string                 `json:"context_url"`
	Contributors                  []*Users               `json:"contributors"`
	Conversation                  interface{}            `json:"conversation"` // Object not documented
	CoverImage                    *Image                 `json:"cover_image"`
	CurrentRevision               *Revisions             `json:"current_revision"`
	CustomTagScope                *CustomTagMessageScope `json:"custom_tag_scope"`
	CustomTags                    []*CustomTags          `json:"custom_tags"`
	Depth                         int                    `json:"depth"`
	DeviceId                      string                 `json:"device_id"`
	ExcludedFromKudosLeaderboards bool                   `json:"excluded_from_kudos_leaderboards"`
	Href                          string                 `json:"href"`
	Id                            string                 `json:"id"`
	Images                        []*Image               `json:"images"`
	IsAnswer                      bool                   `json:"is_answer"`
	IsEscalated                   bool                   `json:"is_escalated"`
	IsImageComment                bool                   `json:"is_image_comment"`
	IsPromoted                    bool                   `json:"is_promoted"`
	IsSolution                    bool                   `json:"is_solution"`
	Kudos                         []*Kudos               `json:"kudos"`
	Labels                        []*Labels              `json:"labels"`
	Language                      string                 `json:"language"`
	LastPublishTime               string                 `json:"last_publish_time"` // JSON DateTime
	ManualSort                    string                 `json:"manual_sort"`
	MeToos                        []*MeToos              `json:"me_toos"`
	Metrics                       interface{}            `json:"metrics"`
	ModerationApprovalDate        string                 `json:"moderation_approval_date"` // JSON DateTime
	ModerationStatus              string                 `json:"moderation_status"`        // string enum
	ModerationStyle               string                 `json:"moderation_style"`         // Hopefully a string enum
	OccasionData                  *OccasionData          `json:"occasion_data"`
	Parent                        interface{}            `json:"parent"`      // Object not documented
	Placeholder                   interface{}            `json:"placeholder"` // Object not documented
	Popularity                    float32                `json:"popularity"`
	PostTime                      string                 `json:"post_time"` // JSON DateTime
	PostTimeFriendly              string                 `json:"post_time_friendly"`
	Ratings                       *Ratings               `json:"ratings"`
	ReadOnly                      bool                   `json:"read_only"`
	Replies                       []*Messages            `json:"replies"`
	SearchSnippet                 string                 `json:"search_snippet"`
	SEODescription                string                 `json:"seo_description"`
	SEOTitle                      string                 `json:"seo_title"`
	SolutionData                  *SolutionsData         `json:"solution_data"`
	Status                        *MessageStatus         `json:"status"`
	Subject                       string                 `json:"subject"`
	Tags                          []*Tags                `json:"tags"`
	Teaser                        string                 `json:"teaser"`
	TKBHelpfulnessRatings         *TKBHelpfulnessRatings `json:"tkb_helpfulness_ratings"`
	Topic                         *Messages              `json:"topics"`
	UserContext                   *UserContext           `json:"user_context"`
	Videos                        []*Videos              `json:"videos"`
	ViewHref                      string                 `json:"view_href"`
	Winner                        bool                   `json:"winner"`
}

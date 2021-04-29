package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/content_workflow
// ContentWorkflow refers to the properties from the returned `content_workflow` object
type ContentWorkflow struct {
	State               string                  `json:"state"` // string enum
	MessageDraftHistory *MessageDraftHistory    `json:"message_draft_history"`
	UserContext         *MessageWorkflowContext `json:"user_context"`
	ViewHref            string                  `json:"view_href"`
}

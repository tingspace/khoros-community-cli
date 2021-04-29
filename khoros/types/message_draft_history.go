package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/message_draft_history
// MessageDraftHistory refers to the properties from the returned `message_draft_history` object
type MessageDraftHistory struct {
	ChangedBy            *Users `json:"changed_by"`
	WorkflowAction       string `json:"workflow_action"`
	ChangedOn            string `json:"changed_on"`             // JSON DateTime
	ScheduledPublishTime string `json:"scheduled_publish_time"` // JSON DateTime
}

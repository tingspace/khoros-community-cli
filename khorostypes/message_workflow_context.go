package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/message_workflow_context
// MessageWorkflowContext refers to the properties from the returned `message_workflow_context` object
type MessageWorkflowContext struct {
	CanSubmitForReview      bool `json:"can_submit_for_review"`
	CanEdit                 bool `json:"can_edit"`
	CanRecall               bool `json:"can_recall"`
	CanSubmitForPublication bool `json:"can_submit_for_publication"`
	CanReturnToAuthor       bool `json:"can_return_to_author"`
	CanPublish              bool `json:"can_publish"`
	CanReturnToReview       bool `json:"can_return_to_review"`
	CanSchedule             bool `json:"can_schedule"`
}

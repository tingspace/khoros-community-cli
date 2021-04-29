package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/user_context
// UserContext refers to the properties from the returned `user_context` object
type UserContext struct {
	CanAcceptSolution    bool `json:"can_accept_solution"`
	CanApprove           bool `json:"can_approve"`
	CanDelete            bool `json:"can_delete"`
	CanEdit              bool `json:"can_edit"`
	CanKudo              bool `json:"can_kudo"`
	CanNominate          bool `json:"can_nominate"`
	CanPromote           bool `json:"can_promote"`
	CanReject            bool `json:"can_reject"`
	CanRejectSolution    bool `json:"can_reject_solution"`
	CanReply             bool `json:"can_reply"`
	CanTag               bool `json:"can_tag"`
	CanSchedule          bool `json:"can_schedule"`
	IsSubscribed         bool `json:"is_subscribed"`
	Kudo                 bool `json:"kudo"`
	MeToo                bool `json:"me_too"`
	Read                 bool `json:"read"`
	TKBHelpfulnessRating bool `json:"tkb_helpfulness_rating"`
}

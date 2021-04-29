package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/occasion_invites
// OccasionInvites refers to the properties from the returned `occasion_invites` object
type OccasionInvites struct {
	Id              string `json:"id"`
	OccasionInvitee *Users `json:"occasion_invitee"`
	InviteMessage   string `json:"invite_message"`
}

package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/membership_requests
// MembershipRequests refers to the properties from the returned `membership_requests` object
type MembershipRequests struct {
	RequestToJoinDate string `json:"request_to_join_date"` // JSON DateTime
}

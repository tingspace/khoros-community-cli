package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/notification_feeds
// NotificationFeeds refers to the properties from the returned `notification_feeds` object
type NotificationFeeds struct {
	Badge             *Badges        `json:"badge"`
	Board             *Boards        `json:"board"`
	Date              string         `json:"date"` // JSON DateTime
	Id                string         `json:"id"`
	Kudos             *Kudos         `json:"kudos"`
	Mentions          interface{}    `json:"mentions"` // Object not documented
	Message           *Messages      `json:"message"`
	NotificationStyle string         `json:"notification_style"` // string enum
	Rank              *Ranks         `json:"rank"`
	Solutions         *SolutionsData `json:"solutions"`
	SubscriptionType  string         `json:"subscription_type"`
	TKB               interface{}    `json:"tkb"` // Object not documented
	Topic             *Messages      `json:"topic"`
}

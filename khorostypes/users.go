package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/users
// Users refers to the properties from the returned `users` object
type Users struct {
	Albums              []*Albums         `json:"albums"`
	Avatar              interface{}       `json:"avatar"` // Object not documented >:(
	Banned              bool              `json:"banned"`
	Biography           string            `json:"biography"`
	BonusPoints         int               `json:"bonus_points"`
	CoverImage          string            `json:"cover_image"`
	Deleted             bool              `json:"deleted"`
	Email               string            `json:"email"`
	EmailExcluded       bool              `json:"email_excluded"`
	FirstName           string            `json:"first_name"`
	Followers           []*Users          `json:"followers"`
	Following           []*Users          `json:"following"`
	Href                string            `json:"href"`
	Id                  string            `json:"id"`
	Images              []*Image          `json:"images"`
	IsOwner             bool              `json:"is_owner"`
	JoinDate            string            `json:"join_date"` // JSON DateTime
	KudosGiven          []*Kudos          `json:"kudos_given"`
	KudosReceived       []*Kudos          `json:"kudos_received"`
	KudosWeight         int               `json:"kudos_weight"`
	Language            string            `json:"language"`
	LastName            string            `json:"last_name"`
	LastVisitTime       string            `json:"last_visit_time"` // JSON DateTime
	Location            string            `json:"location"`
	Login               string            `json:"login"` // Username
	Messages            []*Messages       `json:"messages"`
	Nodes               []*Nodes          `json:"nodes"`
	OnlineStatus        string            `json:"online_status"` // Hoping this is a string enum
	PersonalData        interface{}       `json:"personal_data"` // Object not documented >:(
	PublicImages        []*Image          `json:"public_images"`
	Rank                *Ranks            `json:"rank"`
	RegistrationData    *RegistrationData `json:"registration_data"`
	Reviews             []*Reviews        `json:"reviews"`
	Roles               []*Roles          `json:"roles"`
	Settings            interface{}       `json:"settings"` // Object not documented
	SolutionsAuthored   []*Messages       `json:"solutions_authored"`
	SSOID               string            `json:"sso_id"`
	ThreadsParticipated []*Messages       `json:"threads_participated"`
	Topics              []*Messages       `json:"topics"`
	UserBadges          interface{}       `json:"user_badges"` // Object not documented
	Videos              []*Videos         `json:"videos"`
	ViewHref            string            `json:"view_href"`
	WebPageUrl          string            `json:"web_page_url"`
}

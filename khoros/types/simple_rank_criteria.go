package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/simple_rank_criteria
// SimpleRankCriteria refers to the properties from the returned `simple_rank_criteria` object
type SimpleRankCriteria struct {
	MinutesOnline           int    `json:"minutes_online"`
	MinutesSinceRegistering int    `json:"minutes_since_registering"`
	NumberOfPageViews       int    `json:"number_of_page_views"`
	NumberOfPosts           int    `json:"number_of_posts"`
	NumberOfSignins         int    `json:"number_of_signins"`
	Role                    *Roles `json:"role"`
	SimpleFormula           string `json:"simple_formula"`
	User                    *Users `json:"user"`
}

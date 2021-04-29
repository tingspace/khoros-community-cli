package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/ranks
// Ranks refers to the properties from the returned `ranks` object
type Ranks struct {
	Bold           bool        `json:"bold"`
	Color          string      `json:"color"`
	Formula        string      `json:"formula"`
	FormulaEnabled bool        `json:"formula_enabled"`
	IconLeft       string      `json:"icon_left"`
	IconTopic      string      `json:"icon_topic"`
	Id             string      `json:"id"`
	KudosWeight    int         `json:"kudos_weight"`
	Name           string      `json:"name"`
	Position       int         `json:"position"`
	RankStatus     string      `json:"rank_status"` // Seems to be a string enum
	RolesGranted   []*Roles    `json:"roles_granted"`
	RolesRemoved   []*Roles    `json:"roles_removed"`
	SimpleCriteria interface{} `json:"simple_criteria"` // Object not documented >:(
}

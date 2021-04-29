package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/review_dimension_ratings-1
// ReviewDimensionRatings refers to the properties from the returned `review_dimension_ratings` object
type ReviewDimensionRatings struct {
	Id     string `json:"id"`
	Rating string `json:"rating"`
}

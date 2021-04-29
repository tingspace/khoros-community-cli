package khoros

// https://developer.khoros.com/khoroscommunitydevdocs/docs/product_categories
// ProductCategories refers to the properties from the returned `product_categories` object
type ProductCategories struct {
	Board    *Boards            `json:"board"`
	Id       string             `json:"id"`
	Parent   *ProductCategories `json:"parent"`
	Products []*Products        `json:"products"`
	Status   string             `json:"status"` // string enum
	Title    string             `json:"title"`
}

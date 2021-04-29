package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/products
// Products refers to the properties from the returned `products` object
type Products struct {
	Description       string               `json:"description"`
	Id                string               `json:"id"`
	ImageUrl          string               `json:"image_url"`
	Manufacturer      string               `json:"manufacturer"`
	Price             string               `json:"price"`
	ProductCategories []*ProductCategories `json:"product_categories"`
	ProductUrl        string               `json:"product_url"`
	Status            string               `json:"status"` // string enum
	TagScope          *TaggedProductScope  `json:"tag_scope"`
	Title             string               `json:"title"`
}

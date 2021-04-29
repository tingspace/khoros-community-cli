package khorostypes

// https://developer.khoros.com/khoroscommunitydevdocs/docs/prices
// Prices refers to the properties from the returned `prices` object
type Prices struct {
	Currency interface{} // Docs say boolean, but examples have ints - NFI
	Value    interface{} // Docs say boolean, but examples have strings - NFI
}

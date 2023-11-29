package modules

// Recipe represents a recipe in the app
type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Ingredients  []string `json:"ingredients"`
	Instructions string   `json:"instructions"`
}

// Recipes is a slice to store recipes
var Recipes []Recipe

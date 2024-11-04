package models

// Food represents a food item.
// @Description Food item details
type Food struct {
	// @Description Name of the food
	// @example Burger
	Name string `firestore:"name"`

	// @Description Price of the food
	// @example 10
	Price int `firestore:"price"`

	// @Description Description of the food
	// @example "Delicious beef burger"
	Description string `firestore:"description"`

	// @Description Class of the food (e.g., Main Course, Dessert)
	// @example "Main Course"
	Class string `firestore:"class"`

	// @Description Indicates if the food is disabled
	// @example false
	IsDisabled bool `firestore:"is_disabled"`

	// @Description Supply time for the food
	// @example "30 minutes"
	SupplyTime string `firestore:"supply_time"`
}

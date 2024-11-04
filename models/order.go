package models

type Order struct {
	Id           string         `firestore:"id"`
	RestaurantId string         `firestore:"restaurant_id"`
	UserId       string         `firestore:"user_id"`
	OrderItems   map[string]int `firestore:"order_items"`
	OrderTime    string         `firestore:"order_time"`
	IsFinished   bool           `firestore:"is_finished"`
}

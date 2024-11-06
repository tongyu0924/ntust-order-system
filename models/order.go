package models

type Order struct {
	ID           string         `firestore:"id"` // 主鍵
	RestaurantId string         `firestore:"restaurant_id"`
	UserId       string         `firestore:"user_id"`
	OrderItems   map[string]int `firestore:"order_items"`
	OrderTime    string         `firestore:"order_time"`
	IsFinished   bool           `firestore:"is_finished"`
}

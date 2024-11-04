package models

type Rate struct {
	RestaurantId string `firestore:"restaurant_id"`
	UserId       string `firestore:"user_id"`
	Rate         int    `firestore:"rate"`
	Comment      string `firestore:"comment"`
}

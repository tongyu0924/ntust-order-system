package models

type Menu struct {
	RestaurantId string `firestore:"restaurant_id"`
	Foods        []Food `firestore:"foods"`
}



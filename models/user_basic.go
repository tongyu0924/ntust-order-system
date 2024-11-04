package models

type UserBasic struct {
	ID                string `firestore:"id"` // 主鍵
	Email             string `firestore:"email"`
	Password          string `firestore:"password"`
	Name              string `firestore:"name"`
	Phone             string `firestore:"phone"`
	Roles             string `firestore:"roles"`
	RestaurantId      string `firestore:"restaurant_id"`
	VerificationToken string `firestore:"verification_token,omitempty"`
	IsDisabled        bool   `firestore:"is_disabled"`
}

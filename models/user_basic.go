package models

type UserBasic struct {
	ID                string `firestore:"id"` // 主鍵
	Email             string `firestore:"email" json:"email"`
	Password          string `firestore:"password" json:"password"`
	Name              string `firestore:"name" json:"name"`
	Phone             string `firestore:"phone" json:"phone"`
	Roles             string `firestore:"roles" json:"roles"`
	RestaurantId      string `firestore:"restaurant_id" json:"restaurant_id"`
	VerificationToken string `firestore:"verification_token,omitempty" json:"verification_token"`
	IsDisabled        bool   `firestore:"is_disabled" json:"is_disabled"`
}

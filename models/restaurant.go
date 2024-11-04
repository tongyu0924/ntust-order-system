package models

type Restaurant struct {
	Id         string `firestore:"id"`
	Name       string `firestore:"name"`
	Phone      string `firestore:"phone"`
	Address    string `firestore:"address"`
	OpenTime   string `firestore:"open_time"`
	CloseTime  string `firestore:"close_time"`
	Notice     string `firestore:"notice"`
	IsOpen     bool   `firestore:"is_open"`
	Rate       int    `firestore:"rate"`
	IsDisabled bool   `firestore:"is_disabled"`
}

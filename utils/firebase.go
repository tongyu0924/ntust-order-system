package utils

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

// InitializeFirebase 初始化 Firebase 憑證
func InitializeFirebase() error {
	opt := option.WithCredentialsFile("config/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	FirestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		return err
	}

	log.Println("Firebase initialized successfully.")
	return nil
}

func FindUserByEmail(email string) (*firestore.DocumentSnapshot, error) {
	iter := FirestoreClient.Collection("users").Where("email", "==", email).Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		return doc, nil
	}

	return nil, nil
}

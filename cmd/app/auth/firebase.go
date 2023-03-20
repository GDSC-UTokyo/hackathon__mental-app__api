package auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
)

type FirebaseApp struct {
	*firebase.App
}

var app *FirebaseApp

func init() {
	var err error
	app, err = InitFirebaseApp()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}

func InitFirebaseApp() (*FirebaseApp, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirebaseApp{app}, nil
}

func (app *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

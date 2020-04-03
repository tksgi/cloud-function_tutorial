package helloworld

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"log"
	"os"
)

var projectID = os.Getenv("GCP_PROJECT")
var client *firestore.Client
var ctx context.Context

func init() {
	conf := &firebase.Config{ProjectID: projectID}

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

package resource

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

type FirebaseServices struct {
	App       *firebase.App
	Firestore *firestore.Client
	Auth      *auth.Client
	Storage   *storage.Client
}

func InitializeFirebaseServices(ctx context.Context) (*FirebaseServices, error) {
	jsonFilePath := fmt.Sprintf("./docs/%s", os.Getenv("FB_ADMIN_SA"))
	sa := option.WithCredentialsFile(jsonFilePath)
	projectId := os.Getenv("FB_PROJ_ID")
	fmt.Println("AM i here?")

	config := &firebase.Config{
		ProjectID: projectId,
	}

	// Init App
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		fmt.Println("Is err happening here?")
		return nil, err
	}

	// Init Firestore
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer firestoreClient.Close()

	// Init Auth
	authClient, err := app.Auth(ctx)
	if err != nil {
		firestoreClient.Close()
		return nil, err
	}

	// Init Storage
	storageClient, err := app.Storage(ctx)
	if err != nil {
		firestoreClient.Close()
		return nil, err
	}

	return &FirebaseServices{
		App:       app,
		Firestore: firestoreClient,
		Auth:      authClient,
		Storage:   storageClient,
	}, nil
}

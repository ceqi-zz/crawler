package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatedbClient() (client *mongo.Client, ctx context.Context, err error) {

	uri := "mongodb+srv://cluster0.tmbdu.mongodb.net/crawlerdb?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=/Users/ce/dev/certs/X509-cert-4795979462624466445.pem"
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		client.Disconnect(ctx)
	}

	return
}

package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	//use Ping method to ensure that your MongoDB server was found and connected successfully.
	//Added inside the init function : only called once, like allows to setup db connections,
	//register with various service rgistries, or perform any number of other tasks that you typically only want to do once.
	//"init" are not directly called by the program, they are automatically executed by the Go runtime.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	//Code to create db called "golangApi" and collectionName stores the tasks we will be creating.
	//You also set up collection as a package-level variable so you can reuse the database connection throughout the package.
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}

package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {

	// Define as configurações de conexão
	clientOptions := options.Client().ApplyURI(EnvMongoURI())

	// Conecta ao Mongo
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica conexão ok
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	return client
}

// Instancia o cliente
var DB *mongo.Client = ConnectToDB()

// conseguindo a coleção da db
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("users").Collection(collectionName)
	return collection
}

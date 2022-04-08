package mongodb

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func New(username string, password string, cluster string, database string) (*MongoDB, error) {

	cl, err := getMongoDbconn(username, password, cluster)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := getMongoDbDatabase(cl, database)

	mdb := &MongoDB{
		Client:   cl,
		Database: db,
	}

	return mdb, err
}

func getMongoDbconn(username string, password string, cluster string) (*mongo.Client, error) {

	// uri := "mongodb://" + url.QueryEscape(username) + ":" +
	// 	url.QueryEscape(password) + "@" + cluster + "/" + database
	// print(uri)

	uri_2 := "mongodb://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster + ":27017/"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri_2))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func getMongoDbDatabase(client *mongo.Client, database string) *mongo.Database {
	collection := client.Database(database)
	return collection
}

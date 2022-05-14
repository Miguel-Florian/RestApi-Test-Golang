package config

//go get go.mongodb.org/mongo-driver/mongo
//dep ensure -add "go.mongodb.org/mongo-driver/mongo"
/*import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
}()*/

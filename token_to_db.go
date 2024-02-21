package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func writeRefreshTokenToDB(refreshToken string, createdAt int64, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("auth_tokens").Collection("refresh_tokens")
	_, err = collection.InsertOne(ctx, bson.D{{"refresh_token", refreshToken}, {"created_at", createdAt}, {"user_id", userID}})
	if err != nil {
		return err
	}

	return nil
}

// {{"refresh_token", "created_at", "user_id"}, {refreshToken, createdAt, userID}}

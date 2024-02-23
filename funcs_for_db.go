package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDBCollection() (context.Context, context.CancelFunc, *mongo.Collection, *mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	dbName := os.Getenv("DATABASE_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	collection := client.Database(dbName).Collection(collectionName)

	if err != nil {
		return nil, nil, nil, nil, err
	}
	return ctx, cancel, collection, client, nil
}

func writeRefreshTokenToDB(refreshToken string, validUntil int64, userID string) error {
	ctx, cancel, collection, client, err := connectToDBCollection()
	defer cancel()
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	_, err = collection.InsertOne(ctx, bson.D{{"refresh_token", refreshToken}, {"valid_until", validUntil}, {"user_id", userID}})
	if err != nil {
		return err
	}

	return nil
}

func isUniqueUserId(userID string) error {

	ctx, cancel, collection, client, err := connectToDBCollection()
	defer cancel()
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	docWithSameUserIDCount, err := collection.CountDocuments(ctx, bson.D{{"user_id", userID}})

	if err != nil {
		return err
	}
	if docWithSameUserIDCount > 0 {
		return errors.New("This user already has refresh token")
	}
	return nil
}

func getDocByUserID(userID string) (bson.D, error) {
	ctx, cancel, collection, client, err := connectToDBCollection()
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	var doc bson.D
	err = collection.FindOne(ctx, bson.D{{"user_id", userID}}).Decode(&doc)
	fmt.Println("user_id", userID)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func updateRefreshTokenInDB(userID string, refreshToken string, validUntil int64) error {
	ctx, cancel, collection, client, err := connectToDBCollection()
	defer cancel()
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	_, err = collection.UpdateOne(ctx, bson.D{{"user_id", userID}}, bson.D{{"$set", bson.D{{"refresh_token", refreshToken}, {"valid_until", validUntil}}}})
	if err != nil {
		return err
	}

	return nil
}

// {{"refresh_token", "created_at", "user_id"}, {refreshToken, createdAt, userID}}

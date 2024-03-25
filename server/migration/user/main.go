package main

import (
	"context"
	"fmt"
	"mainframe/user/model"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	userDumpPath := os.Args[1]
	userMigrPath := os.Args[2]

	f, err := os.Open(userDumpPath)
	if err != nil {
		fmt.Printf("Error while opening dump file: %v", err)
		return
	}
	defer f.Close()

	os.Remove(userMigrPath)

	g, err := os.Create(userMigrPath)
	if err != nil {
		fmt.Printf("Error while creating id file: %v", err)
		return
	}
	defer g.Close()

	buffer := make([]byte, 110)

	numberOfBytes, err := f.Read(buffer)
	if err != nil {
		fmt.Printf("Error while reading user file: %v", err)
		return
	}

	for numberOfBytes > 0 {
		row := string(buffer)
		oldId := row[:9]
		username := strings.TrimSpace(row[9:39])
		password := strings.TrimSpace(row[39:59])
		name := strings.TrimSpace(row[59:79])
		surname := strings.TrimSpace(row[79:99])
		birth := strings.TrimSpace(row[99:109])
		newId := primitive.NewObjectID()

		user := model.User{
			Id:       newId,
			Username: username,
			Password: password,
			Name:     name,
			Surname:  surname,
			Birth:    birth,
		}

		fmt.Printf("User with id %v\n", oldId)
		g.Write([]byte(fmt.Sprintf("%v%v\n", oldId, newId.Hex())))

		err = insertUser(user)
		if err != nil {
			fmt.Printf("Error while inserting a user: %v", err)
			return
		}

		numberOfBytes, _ = f.Read(buffer)
	}
}

func insertUser(user model.User) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Insert of a document
	_, err = coll.InsertOne(context.TODO(), user)

	return err
}

func getCollection(db, coll string) (*mongo.Collection, error) {
	// Timeout setup
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second)
	defer cancel()

	// Connection to the db
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"))
	if err != nil {
		return nil, err
	}

	// Check db status
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Retrieve the collection
	return client.Database(db).Collection(coll), nil
}

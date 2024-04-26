package main

import (
	"context"
	"fmt"
	"time"

	//accdb "mainframe/account/db"
	accmod "mainframe/account/model"
	usemod "mainframe/user/model"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	accDumpPath := os.Args[1]
	accMigrPath := os.Args[2]

	f, err := os.Open(accDumpPath)
	if err != nil {
		fmt.Printf("Error while opening dump file: %v", err)
		return
	}
	defer f.Close()

	os.Remove(accMigrPath)

	g, err := os.Create(accMigrPath)
	if err != nil {
		fmt.Printf("Error while creating iban file: %v", err)
		return
	}
	defer g.Close()

	buffer := make([]byte, 98)

	numberOfBytes, err := f.Read(buffer)
	if err != nil {
		fmt.Printf("Error while reading account file: %v", err)
		return
	}

	lastOldAccountId := ""
	lastNewAccountId := primitive.NilObjectID

	for numberOfBytes > 0 {
		row := string(buffer)
		oldId := row[:9]
		oldIban := strings.TrimSpace(row[9:53])
		owner := strings.TrimSpace(row[53:77])
		cashString := strings.TrimSpace(row[77:97])

		cash, _ := strconv.ParseFloat(cashString, 32)

		newId := primitive.NewObjectID()
		newIban := primitive.NewObjectID()

		account := accmod.Account{
			Id:    newId,
			IBAN:  newIban.Hex(),
			Owner: owner,
			Cash:  float32(cash),
		}

		fmt.Printf("Account with id %v, owner %v, cash %v\n", oldId, owner, cash)

		ownerObject, _ := primitive.ObjectIDFromHex(owner)
		user, err := selectUser(ownerObject)
		if err != nil {
			fmt.Printf("Error while searching user %v: %v", owner, err)
			return
		}
		if user.Id == primitive.NilObjectID {
			fmt.Printf("Owner with id %v not found", owner)
			continue
		}

		g.Write([]byte(fmt.Sprintf("%v%v%v%v\n", oldId, newId.Hex(), oldIban, newIban.Hex())))

		if lastOldAccountId == "" || lastOldAccountId != oldId {
			err = insertAccount(account)
			if err != nil {
				fmt.Printf("Error while inserting a account: %v", err)
				return
			}

			lastOldAccountId = oldId
			lastNewAccountId = account.Id
		}

		err = addAccount(ownerObject, lastNewAccountId)
		if err != nil {
			fmt.Printf("Error while inserting a account on user table: %v", err)
			return
		}

		fmt.Printf("Account %v\n", account)

		numberOfBytes, _ = f.Read(buffer)
	}
}

func insertAccount(account accmod.Account) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "accounts")
	if err != nil {
		return err
	}

	// Insert of a document
	_, err = coll.InsertOne(context.TODO(), account)

	return err
}

func selectUser(id primitive.ObjectID) (*usemod.User, error) {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return nil, err
	}

	// Search for a document
	var user usemod.User
	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return &user, err
}

func addAccount(id primitive.ObjectID, account primitive.ObjectID) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Construction of the DB objects
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"accounts": account}}

	// Insert of a document
	_, err = coll.UpdateOne(context.TODO(), filter, update)

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

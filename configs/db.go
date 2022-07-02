package configs

import (
	"context"
	"fmt"
	"log"
	"project-mngt-golang-graphql-gin/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() *DB {
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
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("projectMngt").Collection(collectionName)
}

func (db *DB) CreateProject(input *model.NewProject) *model.Project {
	collection := colHelper(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return &model.Project{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		OwnerID:     input.OwnerID,
		Name:        input.Name,
		Description: input.Description,
		Status:      model.StatusNotStarted,
	}
}

func (db *DB) CreateOwner(input *model.NewOwner) *model.Owner {
	collection := colHelper(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return &model.Owner{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
}

func (db *DB) GetOwners() []*model.Owner {
	collection := colHelper(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var owners []*model.Owner
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleOwner *model.Owner
		if err = res.Decode(&singleOwner); err != nil {
			log.Fatal(err)
		}
		owners = append(owners, singleOwner)
	}

	return owners
}



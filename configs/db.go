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

func (db *DB) CreateProject(input *model.NewProject) (*model.Project, error) {
	collection := colHelper(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	project := &model.Project{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		OwnerID:     input.OwnerID,
		Name:        input.Name,
		Description: input.Description,
		Status:      model.StatusNotStarted,
	}

	return project, err
}

func (db *DB) CreateOwner(input *model.NewOwner) (*model.Owner, error) {
	collection := colHelper(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	owner := &model.Owner{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	return owner, err
}

func (db *DB) GetOwners() ([]*model.Owner, error) {
	collection := colHelper(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var owners []*model.Owner
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleOwner *model.Owner
		if err = res.Decode(&singleOwner); err != nil {
			log.Fatal(err)
		}
		owners = append(owners, singleOwner)
	}

	return owners, err
}

func (db *DB) GetProjects() ([]*model.Project, error) {
	collection := colHelper(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var projects []*model.Project
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleProject *model.Project
		if err = res.Decode(&singleProject); err != nil {
			log.Fatal(err)
		}
		projects = append(projects, singleProject)
	}

	return projects, err
}

func (db *DB) SingleOwner(ID string) (*model.Owner, error) {
	collection := colHelper(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var owner *model.Owner
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&owner)

	return owner, err
}

func (db *DB) SingleProject(ID string) (*model.Project, error) {
	collection := colHelper(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var project *model.Project
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&project)

	return project, err
}

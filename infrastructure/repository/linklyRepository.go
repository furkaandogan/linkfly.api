package repository

import (
	"context"

	"api.linkfly.com/domain/linkly"
	"api.linkfly.com/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	databaseName   = "linkfly"
	collectionName = "Linkly"
)

type LinklyRepository struct {
	Database mongo.Database
}

func (respository LinklyRepository) Insert(context context.Context, linkly *linkly.Linkly) (bool, error) {
	collection := respository.Database.GetCollection(databaseName, collectionName)
	result, err := collection.InsertOne(context, linkly)
	if err != nil {
		return false, err
	}
	return result.InsertedID != nil, nil
}

func (respository LinklyRepository) Update(context context.Context, linkly *linkly.Linkly) (bool, error) {
	collection := respository.Database.GetCollection(databaseName, collectionName)
	filter := bson.M{
		"_id": linkly.Id,
	}
	result, err := collection.ReplaceOne(context, filter, linkly)
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

func (respository LinklyRepository) IsExists(context context.Context, hash string) (bool, error) {
	collection := respository.Database.GetCollection(databaseName, collectionName)
	filter := bson.M{
		"hash": hash,
	}

	link := linkly.Linkly{}

	if err := collection.FindOne(context, filter).Decode(&link); err != nil {
		return false, err
	}

	return link.Id != "", nil

}
func (respository LinklyRepository) Get(context context.Context, hash string) (*linkly.Linkly, error) {
	collection := respository.Database.GetCollection(databaseName, collectionName)
	filter := bson.M{
		"hash": hash,
	}

	link := &linkly.Linkly{}

	if err := collection.FindOne(context, filter).Decode(&link); err != nil {
		return nil, err
	}

	return link, nil

}
func (respository LinklyRepository) Gets(context context.Context, userId string) (*[]linkly.Linkly, error) {
	collection := respository.Database.GetCollection(databaseName, collectionName)
	filter := bson.M{
		"user.id": userId,
	}

	link := &linkly.Linkly{}

	if err := collection.FindOne(context, filter).Decode(&link); err != nil {
		return nil, err
	}

	return nil, nil

}

func NewLinklyRepository(database mongo.Database) linkly.ILinklyRepository {
	return &LinklyRepository{
		Database: database,
	}
}

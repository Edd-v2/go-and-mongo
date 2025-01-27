package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		Collection: db.Collection("user"),
	}
}

func (s *UserService) Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return s.Collection.InsertOne(ctx, document)
}

func (s *UserService) FindOne(ctx context.Context, filter bson.M) (bson.M, error) {
	var result bson.M
	err := s.Collection.FindOne(ctx, filter).Decode(&result)
	return result, err
}

func (s *UserService) FindAll(ctx context.Context, filter bson.M) ([]bson.M, error) {
	cur, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, cur.Err()
}

func (s *UserService) UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return s.Collection.UpdateOne(ctx, filter, update)
}

func (s *UserService) DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	return s.Collection.DeleteOne(ctx, filter)
}

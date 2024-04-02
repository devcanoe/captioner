package store

import (
	"context"

	"captioner.com.ng/pkg/constants"
	"captioner.com.ng/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	IStore[T any] interface {
		Count(ctx context.Context) (int, error)
		Get(ctx context.Context, filter []constants.Identifier) (*T, error)
		GetAll(ctx context.Context, filter []constants.Identifier) (*[]T, error)
		Create(ctx context.Context, body interface{}) (interface{}, error)
		Update(ctx context.Context, filter []constants.Identifier, body []constants.Identifier) (*T, error)
		Delete(ctx context.Context, filter []constants.Identifier) error
		DeleteAll(ctx context.Context, filter []constants.Identifier) error
	}

	Store[T any] struct {
		client *mongo.Collection
	}
)

func InitStore[T any](client *mongo.Collection) *Store[T] {
	return &Store[T]{
		client: client,
	}
}

func (s *Store[T]) Count(ctx context.Context) (int64, error) {
	result, err := s.client.EstimatedDocumentCount(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *Store[T]) Get(ctx context.Context, filter []constants.Identifier) (*T, error) {
	var result T
	param := utils.BsonInterfacer(filter)

	if err := s.client.FindOne(ctx, param).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *Store[T]) GetAll(ctx context.Context, filter []constants.Identifier) (*[]T, error) {
	var results []T
	param := utils.BsonInterfacer(filter)

	data, err := s.client.Find(ctx, param)
	if err != nil {
		return nil, err
	}

	defer data.Close(ctx)
	for data.Next(ctx) {
		var result T
		if err := data.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (s *Store[T]) Create(ctx context.Context, body interface{}) (interface{}, error) {

	result, err := s.client.InsertOne(ctx, body)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (s *Store[T]) Update(ctx context.Context, filter []constants.Identifier, body []constants.Identifier) (*T, error) {
	var result T
	param := utils.BsonInterfacer(filter)
	params := utils.BsonInterfacer(body)

	if err := s.client.FindOneAndUpdate(ctx, param, bson.M{"$set": params}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Store[T]) Delete(ctx context.Context, filter []constants.Identifier) error {
	param := utils.BsonInterfacer(filter)

	_, err := s.client.DeleteOne(ctx, param)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store[T]) DeleteAll(ctx context.Context, filter []constants.Identifier) error {
	param := utils.BsonInterfacer(filter)

	_, err := s.client.DeleteMany(ctx, param)
	if err != nil {
		return err
	}
	return nil
}

package db

import (
	"context"
	"fmt"
	"github.com/XFroggyX/go-logger"
	"github.com/XFroggyX/my-little-news-magazine/internal/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logger.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("faild to create user %w", err)
	}

	d.logger.Debug("convert insertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	d.logger.Trace(user)
	return "", fmt.Errorf("faild to convert objectid to hex")
}

func (d *db) FindOne(ctx context.Context, id string) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) Update(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logger.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

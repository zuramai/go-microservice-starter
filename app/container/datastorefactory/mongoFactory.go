package datastorefactory

import (
	"context"
	"time"

	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoFactory struct{}

func (mongof *mongoFactory) Build(c container.Container, datastoreconfig *config.DataStoreConfig) (DataStoreInterface, error) {
	key := datastoreconfig.Code

	// Check if connection is cached in container
	if value, found := c.Get(key); found {
		logger.SugarLog.Debug("Found database connection in container key: ", key)
		return value, nil
	}

	connection, ctx, err := buildConnection(datastoreconfig)

	c.Put(key, connection)
	c.Put("mongoCtx", ctx)

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func buildConnection(datastoreconfig *config.DataStoreConfig) (*mongo.Client, context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(datastoreconfig.UrlAddress))

	if err != nil {
		return nil, nil, err
	}

	return client, ctx, nil
}

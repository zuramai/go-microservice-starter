package datastorefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type cacheGrpcFactory struct{}

func (cacheGrpcF *cacheGrpcFactory) Build(c container.Container, datastoreconfig *config.DataStoreConfig) (DataStoreInterface, error) {
	key := datastoreconfig.Code

	// If connection found in the container, then return the stored connection
	// We do this because we don't want connect over and over again every request
	if value, found := c.Get(key); found {
		return value.(*grpc.ClientConn), nil
	}

	conn, err := grpc.Dial(datastoreconfig.UrlAddress, grpc.WithInsecure())

	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	c.Put(key, conn)
	return conn, err
}

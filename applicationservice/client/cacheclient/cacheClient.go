package cacheclient

import (
	"context"

	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/client/cacheclient/generatedclient"
	"google.golang.org/grpc"
)

type CacheDataGrpc struct {
	Conn *grpc.ClientConn
}

func getCacheClient(conn *grpc.ClientConn) generatedclient.CacheServiceClient {
	return generatedclient.NewCacheServiceClient(conn)
}

func (cdg CacheDataGrpc) Get(key string) ([]byte, error) {
	cacheClient := getCacheClient(cdg.Conn)

	request := generatedclient.GetRequest{
		Key: key,
	}
	res, err := cacheClient.Get(context.Background(), &request)

	if err != nil {
		return nil, err
	} else {
		return res.Value, err
	}
}

func (cdg CacheDataGrpc) Store(key string, value byte) error {
	cacheClient := getCacheClient(cdg.Conn)
	ctx := context.Background()
	_, err := cacheClient.Store(ctx, &generatedclient.StoreRequest{Key: key, Value: []byte{value}})

	if err != nil {
		return err
	} else {
		logger.Log.Debug("store called")
	}
	return nil
}

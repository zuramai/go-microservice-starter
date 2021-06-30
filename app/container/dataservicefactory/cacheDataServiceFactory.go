package dataservicefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/container/datastorefactory"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/client/cacheclient"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type cacheDataServiceFactory struct{}

func (cachedatasf *cacheDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("cacheDataServiceFactory")
	dsc := dataConfig.DataStoreConfig

	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	grpcConn := dsi.(*grpc.ClientConn)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdg := cacheclient.CacheDataGrpc{Conn: grpcConn}
	//logger.Log.Debug("udm:", udm.DB)

	return &cdg, nil
}

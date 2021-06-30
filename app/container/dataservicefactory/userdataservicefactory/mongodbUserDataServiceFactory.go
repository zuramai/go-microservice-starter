package userdataservicefactory

import (
	"github.com/pkg/errors"

	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/container/datastorefactory"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/dataservice"
	"github.com/muka-id/service-user/applicationservice/dataservice/userdata/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbUserDataServiceFactory struct{}

func (mongoudsf *mongodbUserDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("sqlUserDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	connection := dsi.(*mongo.Client)
	database := connection.Database(dataConfig.DataStoreConfig.DbName)
	uds := mongodb.UserDataMongo{Connection: connection, DB: database}
	return &uds, nil
}

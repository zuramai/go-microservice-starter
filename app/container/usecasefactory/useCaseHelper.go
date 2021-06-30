package usecasefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/container/dataservicefactory"
	"github.com/muka-id/service-user/applicationservice/dataservice"
	"github.com/pkg/errors"
)

func buildUserData(c container.Container, dc *config.DataConfig) (dataservice.UserDataInterface, error) {
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udi := dsi.(dataservice.UserDataInterface)
	return udi, nil
}

func buildCacheData(c container.Container, dc *config.DataConfig) (dataservice.CacheDataInterface, error) {
	//logger.Log.Debug("uc:", cdc)
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi := dsi.(dataservice.CacheDataInterface)
	return cdi, nil
}

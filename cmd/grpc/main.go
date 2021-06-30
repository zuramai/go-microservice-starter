package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/muka-id/service-user/app"
	"github.com/muka-id/service-user/app/container/servicecontainer"
	"github.com/muka-id/service-user/app/controller"
	"github.com/muka-id/service-user/app/logger"
	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	DEV_CONFIG  string = "./app/config/app.dev.yaml"
	PROD_CONFIG string = "./app/config/app.prod.yaml"
)

func runServer(sc *servicecontainer.ServiceContainer) error {
	server := grpc.NewServer()
	reflection.Register(server)
	unimplemented := uspb.UnimplementedUserServiceServer{}
	cs := &controller.UserService{sc, unimplemented}
	uspb.RegisterUserServiceServer(server, cs)

	usergc := sc.AppConfig.UserGrpcConfig

	listen, err := net.Listen(usergc.DriverName, usergc.UrlAddress)

	if err != nil {
		return errors.Wrap(err, "")
	} else {
		logger.SugarLog.Debug("Server listening on ", usergc.UrlAddress)
	}

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)

		<-c

		logger.SugarLog.Debug("Stopping the server..")
		server.Stop()
		listen.Close()
		logger.SugarLog.Debug("Done.")

		v, ok := sc.Get(sc.AppConfig.MongoConfig.Code)
		if ok {
			logger.SugarLog.Debug("Closing MongoDB connection..")
			ctx, _ := sc.Get("mongoCtx")
			v.(*mongo.Client).Disconnect(ctx.(context.Context))
		}

	}()

	return server.Serve(listen)

}

func main() {
	filename := DEV_CONFIG

	container, err := buildContainer(filename)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	if err := runServer(container); err != nil {
		logger.SugarLog.Errorf("Failed to run user server: %+v\n", err)
		panic(err)
	}

}

func buildContainer(config string) (*servicecontainer.ServiceContainer, error) {
	container, err := app.InitApp(config)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	sc := container.(*servicecontainer.ServiceContainer)
	return sc, nil
}

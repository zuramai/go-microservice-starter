# Go Microservice Starter
A boilerplate for flexible Go microservice.

# Table of contents
1. [Features](#features)
2. [Installation](#installation)
3. [Todo List](#todo-list)
4. [Folder Structures](#folder-structures)

## Features:
- Multiple database supports
- MongoDB as starter database
- Clean architecture with SOLID Principles
- Factory method implemented
- Separated use case, handler, and data service layer
- gRPC support
- Container caching
- Superfast logging with Zap

## Installation
1. Clone the repository
```
git clone https://github.com/zuramai/go-microservice-starter
cd go microservice-starter
```
2. Install dependencies
```
make install
# OR
go mod download
go mod vendor
```
3. Run server
```
make runserver
```

## Todo list
- [ ] Add ElasticSearch example
- [ ] Add SQL example
- [ ] Add Docker implementation

# Folder structures
This is the top-level project structure:

![image](https://user-images.githubusercontent.com/45036724/123934409-73379180-d9bd-11eb-80aa-063765b907c8.png)

### app
![image](https://user-images.githubusercontent.com/45036724/123934830-dfb29080-d9bd-11eb-8dfc-bd7cc8ab6549.png)

In app folder, it contains `config` folder in which all the service config are located in here. There are two config (yaml) files `app.dev.yaml` for development environment and `app.prod.yaml` for production environment. We also have `appConfig.go` where we map all the config into structs.

### app/container
![image](https://user-images.githubusercontent.com/45036724/123935673-a169a100-d9be-11eb-967f-1ec2d39f5b72.png)

The dependency injection container, which is responsible for creating concrete types and injecting them into each function. We have `containerhelper` folder which is in charge of creating every use case into a concrete type.

The `dataservicefactory` folder is responsible to create concrete type of data service. For instance, we have cacheDataServiceFactory which create connection to Cache GRPC Service. We have userDataServiceFactory which connect to the database, in this context it is MongoDB.

`servicecontainer` is implementation of container interface. It's Only has one file, which is the key for “container” package. The following is the code. The starting point of it is **InitApp**, which reads configuration data from a file and set the logger.

`usecasefactory`: For each use case, such as `registration` , the interface is defined in `usecase` package, but the concrete type is defined in `registration` sub-package under `usecase` package. Also, there is a factory in the container which is responsible to create the concrete use case instance. For the `registration` use case, it is `registrationFactory.go`. The relationship between the use case and the use case factory is one-to-one.


## Reference
Thanks to @jfeng45



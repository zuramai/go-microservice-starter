package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	MongoConfig     DataStoreConfig `yaml:"mongoConfig`
	CacheGrpcConfig DataStoreConfig `yaml:"cacheGrpcConfig"`
	UserGrpcConfig  DataStoreConfig `yaml:"userGrpcConfig"`
	ZapConfig       zap.Config      `yaml:"zapConfig"`
	UseCaseConfig   UseCaseConfig   `yaml:"useCaseConfig"`
}

type UseCaseConfig struct {
	Registration RegistrationConfig `yaml:"registration"`
	ListUser     ListUserConfig     `yaml:"listUser"`
	UpdateUser   UpdateUserConfig   `yaml:"updateUser"`
	DeleteUser   DeleteUserConfig   `yaml:"deleteUser"`
}

// Database code
const (
	MONGODB    string = "mongodb"
	USER_GRPC  string = "userGrpc"
	CACHE_GRPC string = "cacheGrpc"
)

// Use case code
const (
	REGISTRATION string = "registration"
	LIST_USER    string = "listUser"
	DELETE_USER  string = "deleteUser"
	UPDATE_USER  string = "updateUser"
)

// Data service code
const (
	USER_DATA  string = "userData"
	CACHE_DATA string = "cacheData"
)

type ListUserConfig struct {
	Code            string     `yaml:"code"`
	UserDataConfig  DataConfig `yaml:"userDataConfig"`
	CacheDataConfig DataConfig `yaml:"cacheDataConfig"`
}

type RegistrationConfig struct {
	Code           string     `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
}

type DeleteUserConfig struct {
	Code           string     `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
}

type UpdateUserConfig struct {
	Code           string     `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
}

// Data config represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

type DataStoreConfig struct {
	Code       string `yaml:"code"`
	DriverName string `yaml:"driverName"`
	UrlAddress string `yaml:"urlAddress"`
	DbName     string `yaml:"dbName"`
	Tx         bool   `yaml:"tx"`
}

// BuildConfig build the AppConfig
// if the filaname is not empty, then it reads the file of the filename (in the same folder) and put it into the AppConfig
func BuildConfig(filename ...string) (*AppConfig, error) {
	if len(filename) == 1 {
		return buildConfigFromFile(filename[0])
	} else {
		return BuildConfigWithoutFile()
	}
}

// BuildConfigWithoutFile create AppConfig with adhoc value
func BuildConfigWithoutFile() (*AppConfig, error) {
	return nil, nil
}

// buildConfigFromFile reads the file of the filename (in the same folder) and put it into the AppConfig
func buildConfigFromFile(filename string) (*AppConfig, error) {

	var ac AppConfig
	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	file, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}
	return &ac, nil
}

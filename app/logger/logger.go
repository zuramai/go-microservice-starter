package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var Log *zap.Logger
var SugarLog *zap.SugaredLogger

func SetLogger(zapConfig zap.Config) error {
	logger, err := zapConfig.Build()
	if err != nil {
		return errors.Wrap(err, "Error setting new logger")
	}

	Log = logger
	SugarLog = logger.Sugar()

	return nil
}

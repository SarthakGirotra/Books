package logger

import "go.uber.org/zap"

//creates sugar logger
func MakeLogger() *zap.SugaredLogger {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar
}

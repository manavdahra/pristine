package providers

import (
	"go.uber.org/zap"
	"sync"
)

var singleton sync.Once
var logger *zap.SugaredLogger
func InitLogger() *zap.SugaredLogger {
	singleton.Do(func() {
		logger = zap.NewExample().Sugar()
	})
	return logger
}

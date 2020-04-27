package zaplogger

import "go.uber.org/zap"

//NewLogger : setup zap logger
func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"error.log",
	}
	return cfg.Build()
}

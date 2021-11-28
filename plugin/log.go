package plugin

import "go.uber.org/zap"

var LOG_PLUGIN_LOAD_STATUS = func(name string, err error, status bool) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Loading plugin",
		zap.String("PluginName", name),
		zap.Bool("status", status),
		zap.Error(err),
	)
}

var LOG_PLUGIN_UNLOAD_STATUS = func(name string, err error, status bool) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Unloading plugin",
		zap.String("PluginName", name),
		zap.Bool("status", status),
		zap.Error(err),
	)
}

var LOG_UNSUPPORTED_OS = func() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Unsupported OS!")
}

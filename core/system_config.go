package core

import (
	"github.com/sirupsen/logrus"
	"time"
	"will/envconfig/config"
	"will/will_tools/logs"
)

var Environment string
var ConfigInited bool
var CoreConfig map[string]interface{}

func initCoreConfig() (bool, func()) {
	if ConfigInited {
		return true, func() {}
	}

	var res bool
	for i := 0; i < 5; i++ {
		res = coreConfig(Environment)
		if res || i > 5 {
			return res, func() {}
		} else {
			time.Sleep(time.Second * 3)
		}
	}
	return res, func() {}
}

func coreConfig(env string) bool {
	if len(env) <= 0 {
		panic("env can not be null")
	}
	if err := fetchCoreConfig(env); err != nil {
		logInfo := logs.TraceFormatter{
			Trace: logrus.Fields{
				"error": err.Error(),
			},
		}
		_ = Log.Error(logInfo)
		return false
	}
	return true
}

func fetchCoreConfig(env string) error {
	setting, err := config.NewSetting(env)
	if err != nil {
		return err
	}
	CoreConfig = make(map[string]interface{})
	err = setting.ReadSection(env, &CoreConfig)
	if err != nil {
		return err
	}
	ConfigInited = true
	return nil
}

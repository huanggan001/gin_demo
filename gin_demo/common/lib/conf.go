package lib

import (
	log2 "gin_demo/common/log"
)

type HttpConf struct {
	Addr           string `mapstructure:"addr"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
}

type SwaggerConf struct {
	Title    string `mapstructure:"title"`
	Desc     string `mapstructure:"desc"`
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
}

type BaseConf struct {
	DebugMode    string         `mapstructure:"debug_mode"`
	TimeLocation string         `mapstructure:"time_location"`
	Log          log2.LogConfig `mapstructure:"log"`
	Base         struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	} `mapstructure:"base"`
	Http    HttpConf    `mapstructure:"http"`
	Swagger SwaggerConf `mapstructure:"swagger"`
}

var ConfBase *BaseConf

func InitBaseConf(path string) error {
	ConfBase = &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}

	if ConfBase.DebugMode == "" {
		if ConfBase.Base.DebugMode != "" {
			ConfBase.DebugMode = ConfBase.Base.DebugMode
		} else {
			ConfBase.DebugMode = "debug"
		}
	}
	if ConfBase.TimeLocation == "" {
		if ConfBase.Base.TimeLocation != "" {
			ConfBase.TimeLocation = ConfBase.Base.TimeLocation
		} else {
			ConfBase.TimeLocation = "Asia/Chongqing"
		}
	}
	if ConfBase.Log.Level == "" {
		ConfBase.Log.Level = "trace"
	}

	//配置日志
	logConf := log2.LogConfig{
		Level: ConfBase.Log.Level,
		FW: log2.ConfFileWriter{
			On:              ConfBase.Log.FW.On,
			LogPath:         ConfBase.Log.FW.LogPath,
			RotateLogPath:   ConfBase.Log.FW.RotateLogPath,
			WfLogPath:       ConfBase.Log.FW.WfLogPath,
			RotateWfLogPath: ConfBase.Log.FW.RotateWfLogPath,
		},
		CW: log2.ConfConsoleWriter{
			On:    ConfBase.Log.CW.On,
			Color: ConfBase.Log.CW.Color,
		},
	}
	if err := log2.SetupDefaultLogWithConf(logConf); err != nil {
		panic(err)
	}
	log2.SetLayout("2006-01-02T15:04:05.000")
	return nil
}

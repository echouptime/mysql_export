package config

import "github.com/spf13/viper"

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

type Web struct {
	Auth struct {
		User     string
		Password string
	}
}

type Logger struct {
	LoggerLevel string `mapstructure:"loggerlevel"`
	Formats     string `mapstructure:"loggerformat"`
	Filename    string `mapstructure:"filename"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	Compress    bool   `mapstructure:"compress"`
}

type Options struct {
	Mysql  Mysql
	Web    Web
	Logger Logger
}

func ParseConfig(path string) (*Options, error) {
	conf := viper.New()
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}
	options := &Options{}
	if err := conf.Unmarshal(&options); err != nil {
		return nil, err
	}
	return options, nil

}

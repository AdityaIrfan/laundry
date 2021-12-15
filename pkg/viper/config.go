package viper

import "github.com/spf13/viper"

type EnvConfig struct {
	FileName string
	FileType string
	Path     string
}

func (e *EnvConfig) ReadConfig() error {
	viper.SetConfigName(e.FileName)
	viper.SetConfigType(e.FileType)
	viper.AddConfigPath(e.Path)
	viper.AutomaticEnv()
	viper.WatchConfig()
	err := viper.ReadInConfig() //  find and read the config file
	if err != nil {
		return err
	}
	return nil
}
